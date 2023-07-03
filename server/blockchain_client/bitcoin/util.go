package bitcoin

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/binary"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
	"github.com/tyler-smith/go-bip32"
)

func keyToECDSA(key *bip32.Key) (*ecdsa.PrivateKey, error) {
	curve := elliptic.P256() // or the appropriate elliptic curve based on your key's curve type

	privateKey := new(ecdsa.PrivateKey)
	privateKey.PublicKey.Curve = curve

	// Decode the Key.Key bytes to obtain the public key coordinates
	x, y := elliptic.Unmarshal(curve, key.Key)
	privateKey.PublicKey.X = x
	privateKey.PublicKey.Y = y

	// Optionally, pad the key bytes to the correct size if needed
	paddedKey := make([]byte, (curve.Params().BitSize+7)/8)
	copy(paddedKey[len(paddedKey)-len(key.Key):], key.Key)
	privateKey.D = new(big.Int).SetBytes(paddedKey)

	return privateKey, nil
}

func ecdsaToBTCEC(key *ecdsa.PrivateKey) (*btcec.PrivateKey, error) {
	// Convert the private key's D field to a byte slice
	keyBytes := key.D.Bytes()

	// Create a new btcec.PrivateKey using the btcec.NewPrivateKey function
	btcKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), keyBytes)

	return btcKey, nil
}
func bipKeyToPrivateKey(bipKey *bip32.Key) (*btcec.PrivateKey, error) {
	privKey, err := keyToECDSA(bipKey)
	if err != nil {
		return nil, err
	}

	// Convert to btcec.PrivateKey type
	btcecPrivKey, _ := ecdsaToBTCEC(privKey)

	return btcecPrivKey, nil
}

func signWithSecp256k1(hash []byte, privateKey *btcec.PrivateKey) ([]byte, error, *btcec.Signature) {
	signature, err := privateKey.Sign(hash)
	if err != nil {
		return nil, err, nil
	}

	// Serialize the signature
	serialized := signature.Serialize()

	return serialized, nil, signature
}

////

func serializeVarInt(b *bytes.Buffer, n int) int {
	if n < 0xfd {
		b.WriteByte(byte(n))
		return 1
	} else if n <= 0xffff {
		b.WriteByte(0xfd)
		binary.Write(b, binary.LittleEndian, uint16(n))
		return 3
	} else if n <= 0xffffffff {
		b.WriteByte(0xfe)
		binary.Write(b, binary.LittleEndian, uint32(n))
		return 5
	} else {
		b.WriteByte(0xff)
		binary.Write(b, binary.LittleEndian, uint64(n))
		return 9
	}
}

func reverseString(s string) string {
	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)
}
