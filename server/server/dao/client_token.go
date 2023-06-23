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
