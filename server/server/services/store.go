package services

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/dao"
	"github.com/sijibomii/cryptopay/server/util"
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
