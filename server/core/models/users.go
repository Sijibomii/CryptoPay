package models

import (
	"time"

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

func (u *User) Insert(conn *utils.PgExecutorAddr, d UserPayload) error {
	return nil
}
func (u *User) Update(conn *utils.PgExecutorAddr, d UserPayload) error {
	return nil
}
func (u *User) Find_by_reset_token(conn *utils.PgExecutorAddr, d UserPayload) error {
	return nil
}
func (u *User) Find_by_email(conn *utils.PgExecutorAddr, d UserPayload) error {
	return nil
}
func (u *User) Find_by_id(conn *utils.PgExecutorAddr, d UserPayload) error {
	return nil
}
func (u *User) Activate(conn *utils.PgExecutorAddr, d UserPayload) error {
	return nil
}
func (u *User) Delete(conn *utils.PgExecutorAddr, d UserPayload) error {
	return nil
}
func (u *User) DeleteExpired(conn *utils.PgExecutorAddr, d UserPayload) error {
	return nil
}
func (u *User) Export(conn *utils.PgExecutorAddr, d UserPayload) error {
	return nil
}
