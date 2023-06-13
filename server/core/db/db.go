package db

import (
	"fmt"

	"github.com/anthdm/hollywood/actor"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/core/utils"
)

type DBClient struct {
	utils.PgExecutor
}

func (d *DBClient) Receive(ctx *actor.Context) {
	switch l := ctx.Message().(type) {
	case actor.Started:
		fmt.Println("User db actor started")

	case models.InsertUserMessage:
		payload, err := insertUser(d.DB, l.Payload)
		if err != nil {
			ctx.Respond(nil)
		}
		ctx.Respond(payload)

	case models.UpdateUserMessage:
		updateUser(d.DB, l.Id, l.Payload)

	case models.FindUserByEmailMessage:
		findUserByEmail(d.DB, l.Email)

	case models.FindUserByIdMessage:
		findUserById(d.DB, l.Id)

	case models.FindUserByResetTokenMessage:
		findUserByResetToken(d.DB, l.Token)

	case models.ActivateUserMessage:
		activateUser(d.DB, l.Token)

	case models.DeleteUserMessage:
		deleteUser(d.DB, l.Token)

	case models.DeleteExpiredUserMessage:
		deleteExpiredUser(d.DB, l.Email)

	// add store cases below

	default:
		fmt.Println("UNKNOWN MESSAGE TO USER DB")
	}
}
