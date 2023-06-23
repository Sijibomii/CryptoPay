package services

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/dao"
	"github.com/sijibomii/cryptopay/server/util"
)

func GetAllClientTokensByStoreId(appState *util.AppState, id uuid.UUID, offset, limit int) ([]models.ClientToken, error) {
	var tokens []models.ClientToken
	var err error

	tokens, err = dao.GetClientTokenListByStoreId(appState.Engine, appState.Postgres, id, offset, limit)

	if err != nil {
		return nil, errors.Wrap(err, "error occurs")
	}

	return tokens, nil
}
