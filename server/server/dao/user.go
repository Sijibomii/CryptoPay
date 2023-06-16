package dao

// interacts with db
import (
	"github.com/anthdm/hollywood/actor"
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

// creat
