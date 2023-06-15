package dao

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/util"
)

func CreateSession(e *actor.Engine, conn *actor.PID, session *models.Session) error {
	_, err := models.InsertSession(e, conn, session)

	if err != nil {
		return util.NewErrNotFound("user not found")
	}

	return nil
}
