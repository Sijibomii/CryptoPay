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
		payload, err := findUserByEmail(d.DB, l.Email)
		fmt.Printf("OUtput of actor %v \n", payload)

		if err != nil {
			ctx.Respond(nil)
		}
		ctx.Respond(payload)

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
	case models.InsertStoreMessage:
		insertStore(d.DB, l.Payload)

	case models.UpdateStoreMessage:
		updateStore(d.DB, l.Id, l.Payload)

	case models.FindStoreByIdMessage:
		findStoreById(d.DB, l.Id)

	case models.FindStoreByOwnerMessage:
		findStoreByOwner(d.DB, l.OwnerID, l.Limit, l.Offset)

	case models.FindStoreByIdWithDeletedMessage:
		findStoreByIdWithDeleted(d.DB, l.Id)

	case models.DeleteStoreMessage:
		deleteStore(d.DB, l.Id)

	case models.SoftDeleteStoreMessage:
		softDeleteStore(d.DB, l.Id)

	case models.SoftDeleteStoreByOwnerIDMessage:
		softDeleteStoreByOwnerID(d.DB, l.OwnerID)

	// client tokens
	case models.InsertClientTokenMessage:
		insertClientToken(d.DB, l.Payload)

	case models.FindClientTokensByStoreMessage:
		findClientTokensByStore(d.DB, l.Store_id, l.Limit, l.Offset)

	case models.FindClientTokenByIdMessage:
		findClientTokenById(d.DB, l.Id)

	case models.FindClientTokenByTokenAndDomainMessage:
		findClientTokenByTokenAndDomain(d.DB, l.Token, l.Domain)

	case models.DeleteClientTokenMessage:
		deleteClientToken(d.DB, l.Id)

	// sesssion

	case models.InsertSessionMessage:
		// panic if there's error
		payload := insertSession(d.DB, l.Payload)

		// how to catch a panic?
		ctx.Respond(payload)

	default:
		fmt.Println("UNKNOWN MESSAGE TO USER DB")
	}
}
