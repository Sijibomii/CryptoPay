package dao

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/google/uuid"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/util"
)

func GetClientTokenListByStoreId(e *actor.Engine, conn *actor.PID, id uuid.UUID, offset, limit int) ([]models.ClientToken, error) {
	tokens, err := models.FindClientTokensByStore(e, conn, id, int64(limit), int64(offset))

	if err != nil {
		return nil, util.NewErrNotFound("store not found")
	}

	return tokens, nil
}

func CreateClientToken(e *actor.Engine, conn *actor.PID, store_id uuid.UUID, name, domain string) (*models.ClientToken, error) {
	token, err := models.InsertClientToken(e, conn, models.ClientTokenPayload{
		ID:       uuid.New(),
		Name:     name,
		Token:    uuid.New(),
		Store_id: store_id,
		Domain:   domain,
	})
	if err != nil {
		return nil, util.NewErrNotFound("creation failed")
	}

	return &token, nil
}

func GetClientTokenById(e *actor.Engine, conn *actor.PID, id uuid.UUID) (*models.ClientToken, error) {
	token, err := models.FindClientTokenById(e, conn, id)
	if err != nil {
		return nil, util.NewErrNotFound("get client token by id failed")
	}

	return &token, nil
}
