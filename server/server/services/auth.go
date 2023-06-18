package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/dao"
	"github.com/sijibomii/cryptopay/server/util"
)

func Login(appState *util.AppState, email string, password string) (string, error) {
	var user *models.User

	if email != "" {
		var err error
		user, err = dao.GetUserByEmail(appState.Engine, appState.Postgres, email)
		if err != nil && !util.IsErrNotFound(err) {
			return "", errors.Wrap(err, "invalid username or password")
		}
	}

	if user == nil {
		return "", errors.New("invalid username or password")
	}

	if !util.ComparePassword(user.Password, password) {
		fmt.Printf("Invalid password for user, userID: %s", user.ID)
		return "", errors.New("invalid username or password")
	}

	session := models.Session{
		ID:     util.NewID(util.IDTypeSession),
		Token:  util.NewID(util.IDTypeToken),
		UserID: user.ID.String(),
		Props:  map[string]interface{}{},
	}

	// store session in db and things like that
	err := dao.CreateSession(appState.Engine, appState.Postgres, &session)
	if err != nil {
		return "", errors.Wrap(err, "unable to create session")
	}

	return session.Token, nil
}

func Register(appState *util.AppState, email string, password string) (*models.User, error) {
	var user *models.User

	if email != "" {
		var err error
		_, err = dao.GetUserByEmail(appState.Engine, appState.Postgres, email)
		if err == nil || !util.IsErrNotFound(err) {
			return nil, errors.Wrap(err, "email has been taken")
		}
		err = util.IsPasswordValid(password, util.PasswordSettings{
			MinimumLength: 9,
			Lowercase:     true,
			Number:        true,
			Uppercase:     true,
			Symbol:        true,
		})
		if err != nil {
			return nil, errors.Wrap(err, "Invalid password")
		}

		// register user
		user, err = dao.RegisterUserByEmail(appState.Engine, appState.Postgres, email, util.HashPassword(password))

		if user == nil {
			errors.New("Reistration failed")
		}

		// send mail
		// appState.Engine.Send(appState.Mailer, mailer.SendActivationMailMessage{
		// 	Payload: *user,
		// })

		return user, nil
	}
	return nil, errors.New("invalid username or password")
}

func Activate(appState *util.AppState, token string) error {

	if token != "" {
		var err error
		parsedUUID, err := uuid.Parse(token)
		_, err = dao.FindUserByActivationTokenAndActivate(appState.Engine, appState.Postgres, parsedUUID)
		if err != nil {
			return errors.Wrap(err, "activation failed")
		}
		return nil
	}
	return nil
}

// resend activation email
