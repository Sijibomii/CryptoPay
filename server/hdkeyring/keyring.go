package hdkeyring

import (
	"github.com/sijibomii/cryptopay/hdkeyring/bip39"
	"github.com/sijibomii/cryptopay/types/bitcoin"
	// "github.com/tyler-smith/go-bip32"
	// "github.com/tyler-smith/go-bip39"
)

type HdKeyring struct {
	Mnemonic   bip39.Mnemonic
	HdPath     DerivationPath
	Wallets    []Wallet
	HdWallet   *XKeyPair
	Root       *XKeyPair
	BtcNetwork bitcoin.Network
}

func NewHdKeyring(path string, numberOfAccounts uint32, btcNetwork bitcoin.Network) (*HdKeyring, error) {

	newPath, err := ParseDerivationPath(path)

	if err != nil {
		//
	}

	mnemonic, err := bip39.NewMnemonic(bip39.Words12, bip39.English, "")

	if err != nil {
		//
	}

}

func initfromMnemonic(Mnemonic bip39.Mnemonic, path DerivationPath, btcNetwork bitcoin.Network) (*HdKeyring, error) {
	master_node, _ := FromSeed(Mnemonic.Seed(), btcNetwork)

	root, _ := master_node.FromPath(path)

	return &HdKeyring{
		Mnemonic:   Mnemonic,
		HdPath:     path,
		HdWallet:   &master_node,
		BtcNetwork: btcNetwork,
		Root:       &root,
	}, nil
}
