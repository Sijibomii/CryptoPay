package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/hdkeyring"
	"github.com/sijibomii/cryptopay/server/dao"
	"github.com/sijibomii/cryptopay/server/util"
	"github.com/sijibomii/cryptopay/types/bitcoin"
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
	var err error

	privateKey, publicKey, err := util.GenerateRSA()

	if err != nil {
		fmt.Printf("error genrating rsa")
	}

	keyring, _ := hdkeyring.NewHdKeyring("m/44'/60'/0'/0", 1, bitcoin.Test)

	store, err = dao.CreateStore(appState.Engine, appState.Postgres, ownerId, name, description, privateKey, publicKey, keyring.Mnemonic.Phrase(), keyring.HdPath.String())

	if err != nil {
		return nil, errors.Wrap(err, "error geting store")
	}

	return store, nil
}
