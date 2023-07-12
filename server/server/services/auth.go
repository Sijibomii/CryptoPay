package services

import (
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
		//fmt.Printf("email is %v \n", user)
		if err != nil && !util.IsErrNotFound(err) {
			return "", errors.Wrap(err, "invalid username or password")
		}
	}

	if user == nil {
		return "", errors.New("invalid username or password")
	}

	if !util.ComparePassword(user.Password, password) {
		//fmt.Printf("Invalid password for user, userID: %s", user.ID)
		return "", errors.New("invalid username or password")
	}

	session := models.Session{
		ID:     util.NewID(util.IDTypeSession),
		Token:  util.NewID(util.IDTypeToken),
		UserID: user.ID,
	}

	// store session in db and things like that
	err := dao.CreateSession(appState.Engine, appState.Postgres, &session)
	if err != nil {
		return "", errors.Wrap(err, "unable to create session")
	}

	return session.Token, nil
}

func ResetPassword(appState *util.AppState, email string) (*models.User, error) {
	var user *models.User

	var err error

	if email != "" {
		var err error
		user, err = dao.GetUserByEmail(appState.Engine, appState.Postgres, email)
		//fmt.Printf("email is %v \n", user)
		if err != nil && !util.IsErrNotFound(err) {
			return user, errors.Wrap(err, "invalid username or password")
		}
	}

	if user == nil {
		return user, errors.New("invalid username or password")
	}

	user, err = dao.SetResetPasswordTokenByEmail(appState.Engine, appState.Postgres, user.ID)

	if err != nil {
		return nil, errors.Wrap(err, "error occurs")
	}

	// send mail
	// appState.Engine.Send(appState.Mailer, mailer.SendResetPasswordMailMessage{
	// 	Payload: *user,
	// })

	return user, nil
}

func ChangePassword(appState *util.AppState, token uuid.UUID, newPassword string) (*models.User, error) {
	var user *models.User

	var err error

	user, err = dao.FindUserByRestToken(appState.Engine, appState.Postgres, token)

	if user == nil {
		errors.New("Operation Failed")
	}

	// hash password
	err = util.IsPasswordValid(newPassword, util.PasswordSettings{
		MinimumLength: 9,
		Lowercase:     true,
		Number:        true,
		Uppercase:     true,
		Symbol:        true,
	})

	if err != nil {
		return nil, errors.Wrap(err, "Invalid password")
	}

	user, err = dao.SetNewPasswordById(appState.Engine, appState.Postgres, user.ID, util.HashPassword(newPassword))

	if err != nil {
		return nil, errors.Wrap(err, "Password Change operation failed")
	}

	return user, nil
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
