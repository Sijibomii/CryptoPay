package dao

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/google/uuid"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/util"
	"github.com/tyler-smith/go-bip32"
)

func GetStoreById(e *actor.Engine, conn *actor.PID, id uuid.UUID) (*models.Store, error) {
	store, err := models.Find_Store_By_Id(e, conn, id)

	if err != nil {
		// user not found
		return nil, util.NewErrNotFound("store not found")
	}

	return &store, nil
}

func FindStoresByOwnerId(e *actor.Engine, conn *actor.PID, ownerId uuid.UUID, offset, limit int) ([]models.Store, error) {
	stores, err := models.Find_Store_By_Owner_Id(e, conn, ownerId, int64(limit), int64(offset))
	if err != nil {
		// user not found
		return nil, util.NewErrNotFound("store not found")
	}

	return stores, nil
}

func CreateStore(e *actor.Engine, conn *actor.PID, ownerId uuid.UUID, name, description string, privateKey *bip32.Key, publicKey *bip32.Key, Mnemonic, Hd_path string) (*models.Store, error) {
	store, err := models.InsertStore(e, conn, models.StorePayload{
		Name:                       name,
		Description:                description,
		Owner_id:                   ownerId,
		ID:                         uuid.New(),
		Btc_confirmations_required: 1,
		Mnemonic:                   Mnemonic,
		Hd_path:                    Hd_path,
		Private_key:                *privateKey,
		Public_key:                 *publicKey,
	})

	if err != nil {
		return nil, util.NewErrNotFound("store failed to be created")
	}

	return &store, nil
}
