package services

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/dao"
	"github.com/sijibomii/cryptopay/server/util"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func FindStoreById(appState *util.AppState, id uuid.UUID) (*models.Store, error) {
	var store *models.Store
	var err error

	store, err = dao.GetStoreById(appState.Engine, appState.Postgres, id)

	if err != nil {
		return nil, errors.Wrap(err, "error geting store")
	}

	return store, nil
}

func FindStoresByOwnerId(appState *util.AppState, ownerId uuid.UUID, offset, limit int) ([]models.Store, error) {
	var stores []models.Store
	var err error

	stores, err = dao.FindStoresByOwnerId(appState.Engine, appState.Postgres, ownerId, offset, limit)

	if err != nil {
		return nil, errors.Wrap(err, "error occurs")
	}

	return stores, nil
}

func CreateStore(appState *util.AppState, ownerId uuid.UUID, name, description string) (*models.Store, error) {
	var store *models.Store
	// var err error

	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	seed := bip39.NewSeed(mnemonic, "")

	masterKey, _ := bip32.NewMasterKey(seed)

	childPath := "m/44'/60'/0'/0" // Replace with your desired path
	childKey, _ := util.NewChildKeyFromString(masterKey, childPath)

	publicKey := childKey.PublicKey()

	store, err := dao.CreateStore(appState.Engine, appState.Postgres, ownerId, name, description, childKey, publicKey, mnemonic, childPath)

	if err != nil {
		return nil, errors.Wrap(err, "error geting store")
	}

	return store, nil
}

func UpdateStoresById(appState *util.AppState, storeId uuid.UUID, name, description string) (*models.Store, error) {
	var store *models.Store
	var err error

	store, err = dao.UpdateStoresById(appState.Engine, appState.Postgres, storeId, name, description)

	if err != nil {
		return nil, errors.Wrap(err, "error geting store")
	}

	return store, nil
}
