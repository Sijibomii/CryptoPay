package dao

// interacts with db
import (
	"github.com/anthdm/hollywood/actor"
	"github.com/google/uuid"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/util"
)

func GetUserByEmail(e *actor.Engine, conn *actor.PID, email string) (*models.User, error) {
	user, err := models.Find_by_email(e, conn, email)

	if err != nil {
		// user not found
		return nil, util.NewErrNotFound("user not found")
	}

	return &user, nil
}

func FindUserByRestToken(e *actor.Engine, conn *actor.PID, token uuid.UUID) (*models.User, error) {
	user, err := models.Find_by_reset_token(e, conn, token)

	if err != nil {
		// user not found
		return nil, util.NewErrNotFound("user not found")
	}

	return &user, nil
}

func RegisterUserByEmail(e *actor.Engine, conn *actor.PID, email, password string) (*models.User, error) {
	user, err := models.InsertUser(e, conn, models.UserPayload{
		Email:    email,
		Password: password,
	})

	if err != nil {
		// reg failed
		return nil, util.NewErrUserRegisterationFailed(email)
	}

	return &user, nil
}

func SetNewPasswordById(e *actor.Engine, conn *actor.PID, id uuid.UUID, newPassword string) (*models.User, error) {

	userPayload := models.UserPayload{
		Password: newPassword,
	}

	user, err := models.UpdateUser(e, conn, id, userPayload)

	if err != nil {
		return nil, util.NewErrNotFound("user could not be updated")
	}

	return &user, nil
}

func SetResetPasswordTokenByEmail(e *actor.Engine, conn *actor.PID, id uuid.UUID) (*models.User, error) {

	userPayload := models.UserPayload{}

	userPayload.Set_reset_token()

	user, err := models.UpdateUser(e, conn, id, userPayload)

	if err != nil {
		return nil, util.NewErrNotFound("user could not be updated")
	}

	return &user, nil
}

func FindUserByActivationTokenAndActivate(e *actor.Engine, conn *actor.PID, token uuid.UUID) (*models.User, error) {
	user, err := models.ActivateUser(e, conn, token)
	if err != nil {
		return nil, util.NewErrUserRegisterationFailed("")
	}
	return &user, nil
}
