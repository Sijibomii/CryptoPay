package hdkeyring

import (
	"fmt"

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
		fmt.Printf("error while parsing derivation path")
	}

	mnemonic, err := bip39.NewMnemonic(bip39.Words12, bip39.English, "")

	if err != nil {
		fmt.Printf("error while creating mnemonic")
	}

	keyring, _ := initfromMnemonic(mnemonic, newPath, btcNetwork)

	keyring.loadWallets(int(numberOfAccounts))

	return keyring, nil
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

func (hd HdKeyring) loadWallets(no_of_accounts int) {

	var wallets []Wallet

	for i := int(0); i < no_of_accounts; i++ {

		keypair, err := hd.Root.Derive(Index{
			value:      uint32(i),
			isHardened: false,
		})

		if err != nil {
			fmt.Printf("error loading wallets")
		}

		wallet := NewWallet(keypair.xprv.secretKey, hd.BtcNetwork)

		wallets = append(wallets, *wallet)
	}
}
