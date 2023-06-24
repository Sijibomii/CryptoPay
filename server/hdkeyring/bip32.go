package hdkeyring

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strconv"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
	"github.com/sijibomii/cryptopay/hdkeyring/bip39"
	"github.com/sijibomii/cryptopay/types/bitcoin"
	"golang.org/x/crypto/ripemd160"
)

const (
	MasterSecret    = "Bitcoin seed"
	HardenedOffset  = 0x80000000
	FingerprintSize = 4
	ChainCodeSize   = 32
)

var (
	ErrInvalidDerivationPath = errors.New("invalid derivation path")
	RegexMasterKey           = regexp.MustCompile(`^[mM]{1}$`)
)

type Fingerprint [FingerprintSize]byte

func (fp *Fingerprint) fromBytes(bytes []byte) {
	copy(fp[:], bytes)
}

type ChainCode [ChainCodeSize]byte

func (cc *ChainCode) fromBytes(bytes []byte) {
	copy(cc[:], bytes)
}

type Index struct {
	isHardened bool
	value      uint32
}

func NewHardenedIndex(value uint32) Index {
	return Index{
		isHardened: true,
		value:      value,
	}
}

func NewSoftIndex(value uint32) Index {

	return Index{
		isHardened: false,
		value:      value,
	}
}

type DerivationPath struct {
	indices []Index
}

func NewDerivationPath(indices []Index) DerivationPath {
	return DerivationPath{
		indices: indices,
	}
}

func (dp *DerivationPath) getIndices() []Index {
	return dp.indices
}

func (dp *DerivationPath) String() string {
	path := "m"
	for _, index := range dp.indices {
		path += "/"
		if index.isHardened {
			path += fmt.Sprintf("%d'", index.value-HardenedOffset)
		} else {
			path += strconv.FormatUint(uint64(index.value), 10)
		}
	}
	return path
}

func ParseDerivationPath(path string) (DerivationPath, error) {
	entries := regexp.MustCompile(`/`).Split(path, -1)

	if len(entries) == 0 || !RegexMasterKey.MatchString(entries[0]) {
		return DerivationPath{}, ErrInvalidDerivationPath
	}

	var indices []Index

	for _, entry := range entries[1:] {
		var index Index
		if len(entry) > 1 && entry[len(entry)-1] == '\'' {
			value, err := strconv.ParseUint(entry[:len(entry)-1], 10, 32)
			if err != nil {
				return DerivationPath{}, ErrInvalidDerivationPath
			}
			index = NewHardenedIndex(uint32(value) + HardenedOffset)
		} else {
			value, err := strconv.ParseUint(entry, 10, 32)
			if err != nil {
				return DerivationPath{}, ErrInvalidDerivationPath
			}
			index = NewSoftIndex(uint32(value))
		}
		indices = append(indices, index)
	}

	return NewDerivationPath(indices), nil
}

// Xprv represents the extended private key.
type Xprv struct {
	secretKey         *btcec.PrivateKey
	ChainCode         ChainCode
	network           bitcoin.Network
	depth             uint32
	index             Index
	parentFingerprint Fingerprint
}

// NewXprv creates an Xprv from the master seed and Bitcoin network.
func NewXprv(seed bip39.Seed, btcNetwork bitcoin.Network) (Xprv, error) {
	var xprv Xprv

	// Create HMAC instance using SHA-512 hash function.
	mac := hmac.New(sha512.New, []byte(MasterSecret))
	// Compute HMAC of the seed.
	mac.Write(seed.AsBytes())
	i := mac.Sum(nil)

	// Split the computed HMAC into secret key and chain code.

	// this creates secretKey and publickey
	secretKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), i[:32])

	xprv.secretKey = secretKey
	xprv.ChainCode = ChainCode(i[32:])
	xprv.network = btcNetwork
	xprv.depth = 0
	xprv.index = Index{
		isHardened: false,
		value:      0,
	}
	xprv.parentFingerprint = Fingerprint{}

	return xprv, nil
}

// Derive derives a child extended private key from the current Xprv using the given derivation path.
func (xprv Xprv) Derive(path DerivationPath) (Xprv, error) {
	for _, index := range path.getIndices() {
		var err error
		xprv, err = xprv.ckdPriv(index)
		if err != nil {
			return xprv, err
		}
	}
	return xprv, nil
}

// ckdPriv derives a child extended private key from the current Xprv using the given index.
func (xprv Xprv) ckdPriv(index Index) (Xprv, error) {
	var newKey Xprv

	// Create HMAC instance using SHA-512 hash function and chain code from the current Xprv.
	mac := hmac.New(sha512.New, xprv.ChainCode[:])
	// Create a secp256k1 curve instance
	// secp := btcec.S256()

	if index.isHardened {
		// If index is a hard index, append 0 byte and the serialized secret key to the HMAC input.
		mac.Write([]byte{0})
		mac.Write(xprv.secretKey.Serialize())
	} else {
		// If index is a soft index, derive the public key from the secret key and append its serialization to the HMAC input.
		publicKey := xprv.secretKey.PubKey()
		mac.Write(publicKey.SerializeCompressed())
	}

	// Append the serialized index to the HMAC input.
	indexBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(indexBytes, index.value)
	mac.Write(indexBytes)

	result := mac.Sum(nil)

	// Split the HMAC result into secret key and chain code.

	//the underscore is a public key
	secretKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), result[:32])

	// Compute the child key by adding the secret key of the current Xprv to the derived secret key.
	newKey.secretKey, _ = computeChildKey(newKey.secretKey, secretKey)

	newKey.depth = xprv.depth + 1
	finger, _ := xprv.Fingerprint()
	newKey.parentFingerprint = Fingerprint(finger[:])
	newKey.index = index
	newKey.ChainCode = ChainCode(result[32:])
	newKey.network = xprv.network

	return newKey, nil
}

func computeChildKey(xprv *btcec.PrivateKey, derivedSecretKey *btcec.PrivateKey) (*btcec.PrivateKey, error) {
	// Create a new private key for the child key
	childKey := new(ecdsa.PrivateKey)

	// Add the secret key of the current Xprv to the derived secret key
	childKey.D = new(big.Int).Add(xprv.D, derivedSecretKey.D)
	childKey.PublicKey = ecdsa.PublicKey{
		Curve: xprv.Curve,
		X:     new(big.Int).Mul(derivedSecretKey.PublicKey.X, xprv.Curve.Params().Gx),
		Y:     new(big.Int).Mul(derivedSecretKey.PublicKey.Y, xprv.Curve.Params().Gy),
	}

	x, _ := btcec.PrivKeyFromBytes(btcec.S256(), childKey.D.Bytes())

	return x, nil
}

// Identifier returns the identifier of the current Xprv.
func (xprv Xprv) Identifier() ([]byte, error) {
	// Derive the public key from the private key
	publicKey := xprv.secretKey.PubKey()

	// Serialize the compressed form of the public key
	serialized := publicKey.SerializeCompressed()

	return serialized, nil
}

// Fingerprint returns the fingerprint of the current Xprv.
func (xprv Xprv) Fingerprint() ([]byte, error) {
	// Create a new secp256k1 curve instance

	// Derive the public key from the private key
	publicKey := xprv.secretKey.PubKey()

	// Compute the identifier by hashing the serialized public key
	identifier := btcutil.Hash160(publicKey.SerializeCompressed())

	// Return the first 4 bytes of the identifier as the fingerprint
	return identifier[:4], nil
}

// AsRaw returns the secret key of the current Xprv.
func (xprv Xprv) AsRaw() *btcec.PrivateKey {
	return xprv.secretKey
}

// String returns the string representation of the current Xprv.
func (xprv Xprv) String() string {
	var buf bytes.Buffer

	switch xprv.network {
	case bitcoin.Mainnet:
		buf.Write([]byte{0x04, 0x88, 0xAD, 0xE4})
	case bitcoin.Test:
		buf.Write([]byte{0x04, 0x35, 0x83, 0x94})
	}

	buf.WriteByte(byte(xprv.depth))

	buf.Write(xprv.parentFingerprint[:])

	indexBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(indexBytes, xprv.index.value)
	buf.Write(indexBytes)

	buf.Write(xprv.ChainCode[:])
	buf.WriteByte(0)
	buf.Write(xprv.secretKey.Serialize())

	checksum := Checksum(buf.Bytes())

	buf.Write(checksum[:4])

	return fmt.Sprintf("%x", buf.Bytes())
}

// FromString creates an Xprv from the string representation.
func FromString(input string) (Xprv, error) {
	var xprv Xprv

	bytes, err := hexToBytes(input)
	if err != nil {
		return xprv, err
	}

	if len(bytes) != 78 {
		return xprv, fmt.Errorf("Invalid key length")
	}

	networkBytes := bytes[:4]
	data := bytes[4:78]
	checksum := bytes[78:]

	switch {
	case bytesEqual(networkBytes, []byte{0x04, 0x88, 0xAD, 0xE4}):
		xprv.network = bitcoin.Mainnet
	case bytesEqual(networkBytes, []byte{0x04, 0x35, 0x83, 0x94}):
		xprv.network = bitcoin.Test
	default:
		return xprv, fmt.Errorf("Invalid network")
	}

	xprv.depth = uint32(data[0])
	xprv.parentFingerprint = ByteToFingerprint(data[1:5])

	indexBytes := make([]byte, 4)
	copy(indexBytes, data[5:9])
	index := binary.BigEndian.Uint32(indexBytes)

	if index < HardenedOffset {
		xprv.index = Index{
			isHardened: false,
			value:      index,
		}
	} else {
		xprv.index = Index{
			isHardened: true,
			value:      index,
		}
	}

	xprv.ChainCode = ByteToChainCode(data[13:45])
	secretKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), data[46:78])

	if err != nil {
		return xprv, err
	}
	xprv.secretKey = secretKey

	expectedChecksum := Checksum(bytes[:78])[:4]
	actualChecksum := checksum

	if !bytesEqual(expectedChecksum, actualChecksum) {
		return xprv, fmt.Errorf("Bad checksum")
	}

	return xprv, nil
}

func ByteToFingerprint(b []byte) Fingerprint {
	var fp Fingerprint
	copy(fp[:], b[:FingerprintSize])
	return fp
}

func ByteToChainCode(b []byte) ChainCode {
	var cc ChainCode
	copy(cc[:], b[:ChainCodeSize])
	return cc
}

// hexToBytes converts a hexadecimal string to a byte slice.
func hexToBytes(input string) ([]byte, error) {
	if len(input)%2 != 0 {
		return nil, fmt.Errorf("Invalid hexadecimal string")
	}

	bytes, err := hex.DecodeString(input)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// bytesEqual checks if two byte slices are equal.
func bytesEqual(a, b []byte) bool {
	return bytes.Compare(a, b) == 0
}

// Checksum calculates the checksum of the data.
func Checksum(data []byte) []byte {
	hash := sha256.Sum256(data)
	hash = sha256.Sum256(hash[:])
	return hash[:]
}

// Xpub represents an extended public key.
type Xpub struct {
	PublicKey         *btcec.PublicKey
	ChainCode         ChainCode
	Network           bitcoin.Network
	Depth             uint32
	Index             Index
	ParentFingerprint Fingerprint
}

// FromPrivate creates an Xpub from an Xprv.
func FromPrivate(xprv *Xprv) (*Xpub, error) {
	// serializedXprv := xprv.Serialize()
	extendedKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, err
	}
	extendedKey.PublicKey.Curve = btcec.S256()
	extendedKey.PublicKey.X, extendedKey.PublicKey.Y = btcec.S256().ScalarBaseMult(xprv.secretKey.Serialize())
	xpub := &Xpub{
		PublicKey:         extendedKey.PubKey(),
		ChainCode:         xprv.ChainCode,
		Network:           xprv.network,
		Depth:             xprv.depth,
		Index:             xprv.index,
		ParentFingerprint: xprv.parentFingerprint,
	}
	return xpub, nil
}

// CkdPub derives a child Xpub at the given index.
func (xpub *Xpub) CkdPub(index *Index) (*Xpub, error) {
	if index.isHardened {
		return nil, fmt.Errorf("invalid derivation")
	}

	indexBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(indexBytes, index.value)

	data := append(xpub.PublicKey.SerializeCompressed(), indexBytes...)
	mac := hmac.New(sha512.New, xpub.ChainCode[:])
	mac.Write(data)
	result := mac.Sum(nil)

	secretKeyBytes := result[:32]
	secretKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), secretKeyBytes)
	publicKey := xpub.PublicKey
	publicKey.X, publicKey.Y = btcec.S256().Add(publicKey.X, publicKey.Y, secretKey.PublicKey.X, secretKey.PublicKey.Y)
	childXpub := &Xpub{
		PublicKey:         publicKey,
		ChainCode:         ChainCode(result[32:]),
		Network:           xpub.Network,
		Depth:             xpub.Depth + 1,
		Index:             *index,
		ParentFingerprint: xpub.Fingerprint(),
	}
	return childXpub, nil
}

// Identifier calculates the identifier of the Xpub.
func (xpub *Xpub) Identifier() []byte {
	hash := ripemd160.New()
	hash.Write(xpub.PublicKey.SerializeCompressed())
	return hash.Sum(nil)
}

// Fingerprint calculates the fingerprint of the Xpub.
func (xpub *Xpub) Fingerprint() Fingerprint {
	identifier := xpub.Identifier()
	return Fingerprint(identifier[:4])
}

// Serialize converts the Xpub to a base58-encoded string.
func (xpub *Xpub) Serialize() string {
	var data []byte
	if xpub.Network == bitcoin.Mainnet {
		data = []byte{0x04, 0x88, 0xB2, 0x1E}
	} else if xpub.Network == bitcoin.Test {
		data = []byte{0x04, 0x35, 0x87, 0xCF}
	}

	data = append(data, byte(xpub.Depth))
	data = append(data, xpub.ParentFingerprint[:]...)
	indexBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(indexBytes, xpub.Index.value)
	data = append(data, indexBytes...)
	data = append(data, xpub.ChainCode[:]...)
	data = append(data, xpub.PublicKey.SerializeCompressed()...)

	checksum := DoubleSHA256(data)[:4]
	serializedData := append(data, checksum...)
	return base58.Encode(serializedData)
}

// DeserializeXpub converts a base58-encoded string to an Xpub.
func DeserializeXpub(xpubString string) (*Xpub, error) {
	serializedData := base58.Decode(xpubString)
	if len(serializedData) != 78 {
		return nil, fmt.Errorf("invalid key length")
	}

	if !bytes.Equal(serializedData[:4], []byte{0x04, 0x88, 0xB2, 0x1E}) &&
		!bytes.Equal(serializedData[:4], []byte{0x04, 0x35, 0x87, 0xCF}) {
		return nil, fmt.Errorf("invalid network")
	}

	expectedChecksum := DoubleSHA256(serializedData[:len(serializedData)-4])[:4]
	actualChecksum := serializedData[len(serializedData)-4:]
	if !bytes.Equal(expectedChecksum, actualChecksum) {
		return nil, fmt.Errorf("bad checksum")
	}

	depth := uint32(serializedData[4])
	parentFingerprint := Fingerprint(serializedData[5:9])
	index := binary.BigEndian.Uint32(serializedData[9:13])
	chainCode := ChainCode(serializedData[13:45])
	publicKey, err := btcec.ParsePubKey(serializedData[45:], btcec.S256())
	if err != nil {
		return nil, err
	}

	xpub := &Xpub{
		PublicKey: publicKey,
		ChainCode: chainCode,
		Depth:     depth,
		Index: Index{
			// is hardened ?
			value: index,
		},
		ParentFingerprint: parentFingerprint,
	}
	if bytes.Equal(serializedData[:4], []byte{0x04, 0x88, 0xB2, 0x1E}) {
		xpub.Network = bitcoin.Mainnet
	} else {
		xpub.Network = bitcoin.Test
	}

	return xpub, nil
}

// DoubleSHA256 computes the double SHA-256 hash of the input data.
func DoubleSHA256(data []byte) []byte {
	hash1 := sha256.Sum256(data)
	hash2 := sha256.Sum256(hash1[:])
	return hash2[:]
}

type XKeyPair struct {
	// xprv represents an extended private key in the BIP32 hierarchy.
	xprv Xprv
	xpub Xpub
}

func NewXKeyPair(xprv Xprv, xpub Xpub) XKeyPair {
	return XKeyPair{xprv, xpub}
}

func FromSeed(seed bip39.Seed, btcNetwork bitcoin.Network) (XKeyPair, error) {
	// create a master private key
	xprv, err := XprvFromMasterSeed(seed, btcNetwork)
	if err != nil {
		return XKeyPair{}, err
	}

	// create a public key from the master private key
	xpub, err := xprv.ToXpub()
	if err != nil {
		return XKeyPair{}, err
	}

	return NewXKeyPair(xprv, xpub), nil
}

func (kp XKeyPair) FromPath(path DerivationPath) (XKeyPair, error) {
	// return a derived private key
	xprv, err := kp.xprv.Derive(path)
	if err != nil {
		return XKeyPair{}, err
	}

	// return the public key created from the private key
	xpub, err := xprv.ToXpub()
	if err != nil {
		return XKeyPair{}, err
	}

	return NewXKeyPair(xprv, xpub), nil
}

func (kp XKeyPair) Derive(index Index) (XKeyPair, error) {
	// creates a private key child from parent private key kp.xprv at a certain index.
	// the index is used in the HMAC hash that created the key.
	xprv, err := kp.xprv.ckdPriv(index)
	if err != nil {
		return XKeyPair{}, err
	}

	xpub, err := xprv.ToXpub()
	if err != nil {
		return XKeyPair{}, err
	}

	return NewXKeyPair(xprv, xpub), nil
}

func (kp XKeyPair) Xprv() Xprv {
	return kp.xprv
}

func (kp XKeyPair) Xpub() Xpub {
	return kp.xpub
}

func XprvFromMasterSeed(seed bip39.Seed, btcNetwork bitcoin.Network) (Xprv, error) {
	// implement the logic for creating a master private key from a seed
	return Xprv{}, errors.New("Not implemented")
}

// func (xprv Xprv) Derive(path DerivationPath) (Xprv, error) {
// 	// implement the logic for deriving a private key from the given path
// 	return Xprv{}, errors.New("Not implemented")
// }

func (xprv Xprv) ToXpub() (Xpub, error) {
	// implement the logic for converting a private key to a public key
	return Xpub{}, errors.New("Not implemented")
}

// func main() {
// 	// Sample usage
// 	seed := bip39.Seed{0x01, 0x02, 0x03} // Replace with actual seed value
// 	btcNetwork := bitcoin.Network(0)     // Replace with actual network value

// 	kp, err := NewXKeyPair(Xprv{}, Xpub{}).FromSeed(seed, btcNetwork)
// 	if err != nil {
// 		fmt.Println("Failed to create key pair:", err)
// 		return
// 	}

// 	// Use the key pair
// 	fmt.Println("xprv:", kp.Xprv())
// 	fmt.Println("xpub:", kp.Xpub())
// }

// func main() {
// 	// Example usage
// 	xprv := &Xprv{} // Define Xprv instance
// 	xpub, err := FromPrivate(xprv)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Xpub:", xpub.Serialize())

// 	// Deserialize Xpub
// 	xpubString := "xpub..."
// 	deserializedXpub, err := DeserializeXpub(xpubString)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Deserialized Xpub:", deserializedXpub.Serialize())
// }

// func main() {
// 	// Example usage
// 	path := "m/0'/1/2"
// 	derivationPath, err := parseDerivationPath(path)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Derivation Path:", derivationPath.String())
// }
