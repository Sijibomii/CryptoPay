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
		payload, err := updateUser(d.DB, l.Id, l.Payload)
		if err != nil {
			panic(err)
		}
		ctx.Respond(payload)

	case models.FindUserByEmailMessage:
		payload, err := findUserByEmail(d.DB, l.Email)
		fmt.Printf("Output of actor %v \n", payload)

		if err != nil {
			ctx.Respond(nil)
		}
		ctx.Respond(payload)

	case models.FindUserByIdMessage:
		payload := findUserById(d.DB, l.Id)

		ctx.Respond(payload)

	case models.FindUserByResetTokenMessage:
		payload := findUserByResetToken(d.DB, l.Token)

		ctx.Respond(payload)

	case models.ActivateUserMessage:
		payload := activateUser(d.DB, l.Token)

		ctx.Respond(payload)

	case models.DeleteUserMessage:
		payload := deleteUser(d.DB, l.Token)

		ctx.Respond(payload)

	case models.DeleteExpiredUserMessage:
		payload := deleteExpiredUser(d.DB, l.Email)

		ctx.Respond(payload)

	// add store cases below
	case models.InsertStoreMessage:
		payload := insertStore(d.DB, l.Payload)

		ctx.Respond(payload)

	case models.UpdateStoreMessage:
		payload := updateStore(d.DB, l.Id, l.Payload)

		ctx.Respond(payload)

	case models.FindStoreByIdMessage:
		payload := findStoreById(d.DB, l.Id)

		ctx.Respond(payload)

	case models.FindStoreByOwnerMessage:
		payload := findStoreByOwner(d.DB, l.OwnerID, l.Limit, l.Offset)

		ctx.Respond(payload)

	case models.FindStoreByIdWithDeletedMessage:
		payload := findStoreByIdWithDeleted(d.DB, l.Id)

		ctx.Respond(payload)

	case models.DeleteStoreMessage:
		payload := deleteStore(d.DB, l.Id)

		ctx.Respond(payload)

	case models.SoftDeleteStoreMessage:
		payload := softDeleteStore(d.DB, l.Id)

		ctx.Respond(payload)

	case models.SoftDeleteStoreByOwnerIDMessage:
		payload := softDeleteStoreByOwnerID(d.DB, l.OwnerID)

		ctx.Respond(payload)

	// client tokens
	case models.InsertClientTokenMessage:
		payload := insertClientToken(d.DB, l.Payload)

		ctx.Respond(payload)

	case models.FindClientTokensByStoreMessage:
		payload := findClientTokensByStore(d.DB, l.Store_id, l.Limit, l.Offset)

		ctx.Respond(payload)

	case models.FindClientTokenByIdMessage:
		payload := findClientTokenById(d.DB, l.Id)

		ctx.Respond(payload)

	case models.FindClientTokenByTokenAndDomainMessage:
		payload := findClientTokenByTokenAndDomain(d.DB, l.Token, l.Domain)

		ctx.Respond(payload)

	case models.DeleteClientTokenMessage:
		payload := deleteClientToken(d.DB, l.Id)

		ctx.Respond(payload)

	// payment

	case models.InsertPaymentMessage:
		payload := insertPayment(d.DB, l.Payload)

		ctx.Respond(payload)

	case models.FindAllPendingPaymentByAddressesMessage:
		payload := findAllPendingPayementByAddresses(d.DB, l.Address, l.Crypto)

		ctx.Respond(payload)

	case models.UpdatePaymentMessage:

		payload, err := updatePayment(d.DB, l.Id, l.Payload)
		if err != nil {
			panic(err)
		}
		ctx.Respond(payload)

	// sesssion

	case models.InsertSessionMessage:
		// panic if there's error
		payload := insertSession(d.DB, l.Payload)

		// how to catch a panic?
		ctx.Respond(payload)

		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Session insertion failed")
				ctx.Respond(nil)
			}
		}()

	case models.GetSessionByTokenMessage:
		payload := getSessionByToken(d.DB, l.Token)

		// how to catch a panic?
		ctx.Respond(payload)

	// Btc Blockchain status

	case models.InsertBtcBlockChainStatusMessage:
		payload := insertBtcBlockChainStatus(d.DB, l.Payload)

		ctx.Respond(payload)

	case models.FindBtcBlockChainStatusByNetworkMessage:
		payload := findBtcBlockChainStatus(d.DB, l.Network)

		ctx.Respond(payload)

	default:
		fmt.Println("UNKNOWN MESSAGE TO USER DB")
	}
}
