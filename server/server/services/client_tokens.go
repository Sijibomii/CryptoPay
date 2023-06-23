package services

import (
	"fmt"

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

func CreateClientToken(appState *util.AppState, store_id uuid.UUID, name, domain string) (*models.ClientToken, error) {
	var token *models.ClientToken
	var err error

	if name == "" {
		//create random str
		name, err = util.GenerateRandomString(20)

		if err != nil {
			fmt.Printf("unable to generate random str for client token name error: %s", err.Error())
			name = "RANDOM API KEY"
		}
	}

	token, err = dao.CreateClientToken(appState.Engine, appState.Postgres, store_id, name, domain)

	if err != nil {
		return nil, errors.Wrap(err, "error occurs")
	}

	return token, nil
}
