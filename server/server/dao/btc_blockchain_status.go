package dao

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/util"
)

func FindBtcBlockChainStatusByNetwork(e *actor.Engine, conn *actor.PID, network string) (*models.BtcBlockChainStatus, error) {
	status, err := models.FindBtcBlockChainStatusByNetwork(e, conn, network)

	if err != nil {
		return nil, util.NewErrNotFound("payment not found")
	}

	return &status, nil
}
