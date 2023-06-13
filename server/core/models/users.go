package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/google/uuid"
	"github.com/sijibomii/cryptopay/core/utils"
)

type UserPayload struct {
	ID                            uuid.UUID
	Email                         string
	Password                      string
	Salt                          string
	Created_at                    time.Time
	Updated_at                    time.Time
	Is_verified                   bool
	Verification_token            uuid.UUID
	Verification_token_expires_at time.Time
	Reset_token                   uuid.UUID
	Reset_token_expires_at        time.Time
}

func (u *UserPayload) Set_created_at() error {
	u.Created_at = time.Now()
	return nil
}
func (u *UserPayload) Set_updated_at() error {
	u.Updated_at = time.Now()
	return nil
}
func (u *UserPayload) Set_verification_token() error {
	u.Is_verified = false
	u.Verification_token = uuid.New()
	u.Verification_token_expires_at = time.Now().Add(time.Hour * 24)
	return nil
}
func (u *UserPayload) Set_reset_token() error {
	u.Reset_token = uuid.New()
	u.Reset_token_expires_at = time.Now().Add(time.Hour * 24)
	return nil
}

func (u *UserPayload) ToUser() User {
	return User{
		ID:                            u.ID,
		Email:                         u.Email,
		Password:                      u.Password,
		Salt:                          u.Salt,
		Created_at:                    u.Created_at,
		Updated_at:                    u.Updated_at,
		Is_verified:                   u.Is_verified,
		Verification_token:            u.Verification_token,
		Verification_token_expires_at: u.Verification_token_expires_at,
		Reset_token:                   u.Reset_token,
		Reset_token_expires_at:        u.Reset_token_expires_at,
	}
}

type User struct {
	ID                            uuid.UUID
	Email                         string
	Password                      string
	Salt                          string
	Created_at                    time.Time
	Updated_at                    time.Time
	Is_verified                   bool
	Verification_token            uuid.UUID
	Verification_token_expires_at time.Time
	Reset_token                   uuid.UUID
	Reset_token_expires_at        time.Time
}

type Insert struct {
	Payload User
}

type Update struct {
	Payload User
	Id      uuid.UUID
}

type FindByEmail struct {
	Email string
}

type FindById struct {
	Id uuid.UUID
}

type FindByResetToken struct {
	Token uuid.UUID
}

type Activate struct {
	Token uuid.UUID
}
type Delete struct {
	Token uuid.UUID
}

type DeleteExpired struct {
	Email string
}

func (u *User) Insert(e *actor.Engine, conn *utils.PgExecutorAddr, d UserPayload) (User, error) {
	var resp = e.Request(conn, Insert{
		Payload: d.ToUser(),
	}, 500)
	res, err := resp.Result()
	if err != nil {
		return User{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(User)

	if !ok {
		return User{}, errors.New("An error occured!")
	}

	return myStruct, nil
}
func (u *User) Update(e *actor.Engine, conn *utils.PgExecutorAddr, id uuid.UUID, d UserPayload) (User, error) {
	var resp = e.Request(conn, Update{
		Payload: d.ToUser(),
		Id:      id,
	}, 500)

	res, err := resp.Result()
	if err != nil {
		return User{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(User)

	if !ok {
		return User{}, errors.New("An error occured!")
	}

	return myStruct, nil

}
func (u *User) Find_by_reset_token(e *actor.Engine, conn *utils.PgExecutorAddr, token uuid.UUID) (User, error) {
	var resp = e.Request(conn, FindByResetToken{
		Token: token,
	}, 500)

	res, err := resp.Result()
	if err != nil {
		return User{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(User)

	if !ok {
		return User{}, errors.New("An error occured!")
	}

	return myStruct, nil
}
func (u *User) Find_by_email(e *actor.Engine, conn *utils.PgExecutorAddr, email string) (User, error) {
	var resp = e.Request(conn, FindByEmail{
		Email: email,
	}, 500)

	res, err := resp.Result()
	if err != nil {
		return User{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(User)

	if !ok {
		return User{}, errors.New("An error occured!")
	}

	return myStruct, nil
}
func (u *User) Find_by_id(e *actor.Engine, conn *utils.PgExecutorAddr, id uuid.UUID) (User, error) {
	var resp = e.Request(conn, FindById{
		Id: id,
	}, 500)

	res, err := resp.Result()
	if err != nil {
		return User{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(User)

	if !ok {
		return User{}, errors.New("An error occured!")
	}

	return myStruct, nil
}
func (u *User) Activate(e *actor.Engine, conn *utils.PgExecutorAddr, token uuid.UUID) (User, error) {
	var resp = e.Request(conn, Activate{
		Token: token,
	}, 500)

	res, err := resp.Result()
	if err != nil {
		return User{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(User)

	if !ok {
		return User{}, errors.New("An error occured!")
	}

	return myStruct, nil
}
func (u *User) Delete(e *actor.Engine, conn *utils.PgExecutorAddr, token uuid.UUID) (User, error) {
	var resp = e.Request(conn, Delete{
		Token: token,
	}, 500)

	res, err := resp.Result()
	if err != nil {
		return User{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(User)

	if !ok {
		return User{}, errors.New("An error occured!")
	}

	return myStruct, nil
}
func (u *User) DeleteExpired(e *actor.Engine, conn *utils.PgExecutorAddr, email string) (User, error) {
	var resp = e.Request(conn, DeleteExpired{
		Email: email,
	}, 500)

	res, err := resp.Result()
	if err != nil {
		return User{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(User)

	if !ok {
		return User{}, errors.New("An error occured!")
	}

	return myStruct, nil
}
func (u *User) Export() ([]byte, error) {
	data := struct {
		ID        uuid.UUID `json:"id"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}{
		ID:        u.ID,
		Email:     u.Email,
		CreatedAt: u.Created_at,
		UpdatedAt: u.Updated_at,
	}

	return json.Marshal(data)
}
