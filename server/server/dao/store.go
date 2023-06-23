package dao

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/google/uuid"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/util"
)

func GetStoreById(e *actor.Engine, conn *actor.PID, id uuid.UUID) (*models.Store, error) {
	store, err := models.Find_Store_By_Id(e, conn, id)

	if err != nil {
		// user not found
		return nil, util.NewErrNotFound("store not found")
	}

	return &store, nil
}
