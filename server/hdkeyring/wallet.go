package hdkeyring

import (
	"crypto/sha256"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
	"github.com/sijibomii/cryptopay/types/bitcoin"
)

type Wallet struct {
	SecretKey  *btcec.PrivateKey
	PublicKey  *btcec.PublicKey
	BtcNetwork bitcoin.Network
}

func NewWallet(secretKey *btcec.PrivateKey, btcNetwork bitcoin.Network) *Wallet {
	publicKey := secretKey.PubKey()

	return &Wallet{
		SecretKey:  secretKey,
		PublicKey:  publicKey,
		BtcNetwork: btcNetwork,
	}
}

func (w *Wallet) GetAddress(currency Crypto) string {
	switch currency {
	case Btc:
		return w.GetBtcAddress()
	case Eth:
		// return w.GetEthAddress()
		return ""
	default:
		return ""
	}
}

// func (w *Wallet) GetBtcsAddress() string {
// 	publicKeyBytes := crypto.FromECDSAPub(w.PublicKey)

// 	// Get RIPEMD160 hash of the public key.
// 	h160 := crypto.RIPEMD160(publicKeyBytes)

// 	var prefixed []byte
// 	if w.BtcNetwork == Mainnet {
// 		prefixed = append([]byte{0x00}, h160...)
// 	} else if w.BtcNetwork == Testnet {
// 		prefixed = append([]byte{0x6F}, h160...)
// 	}

// 	// Get SHA256 hash of the prefixed RIPEMD160 hash.
// 	h256 := crypto.SHA256(prefixed)
// 	addressBytes := append(prefixed, h256[:4]...)

// 	// Base58 encode the address bytes.
// 	address := base58.Encode(addressBytes)
// 	return address
// }

func doubleSha256(data []byte) []byte {
	hash1 := sha256.Sum256(data)
	hash2 := sha256.Sum256(hash1[:])
	return hash2[:]
}

func (w *Wallet) GetBtcAddress() string {
	// h160 on public key.
	h160 := btcutil.Hash160(w.PublicKey.SerializeCompressed())

	// Add version prefix.
	prefixed := make([]byte, 21)
	switch w.BtcNetwork {
	case bitcoin.Mainnet:
		prefixed[0] = 0
	case bitcoin.Test:
		prefixed[0] = 111
	default:
		return ""
	}
	copy(prefixed[1:], h160[:])

	// h256 on prefixed h160.
	h256 := doubleSha256(prefixed)

	// 25 byte binary Bitcoin Address.
	address := make([]byte, 25)
	copy(address, prefixed[:21])
	copy(address[21:], h256[:4])

	// Base58 string of the address.
	return base58.Encode(address)
}

// func main() {
// 	// Example usage
// 	// Generate a new private key
// 	privateKey, _ := btcec.NewPrivateKey(btcec.S256())

// 	// Get the corresponding public key
// 	publicKey := privateKey.PubKey()

// 	// Get the BTC address
// 	btcAddress, _ := getBTCAddress(publicKey, "Mainnet")

// 	fmt.Println("BTC Address:", btcAddress)
// }

type Crypto int

const (
	Btc Crypto = iota
	Eth
)

// func main() {
// 	// Example usage
// 	secretKey, _ := crypto.GenerateKey()
// 	wallet := NewWallet(secretKey, Mainnet)
// 	fmt.Println("BTC Address:", wallet.GetAddress(Btc))
// 	fmt.Println("ETH Address:", wallet.GetAddress(Eth))
// }
