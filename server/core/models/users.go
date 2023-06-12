package models

import (
	"time"

	"github.com/google/uuid"
)

type UserPayload struct {
	ID                            uuid.UUID
	email                         string
	password                      string
	salt                          string
	created_at                    time.Time
	updated_at                    time.Time
	is_verified                   bool
	verification_token            uuid.UUID
	verification_token_expires_at time.Time
	reset_token                   uuid.UUID
	reset_token_expires_at        time.Time
}

func (u *UserPayload) set_created_at() error {
	u.created_at = time.Now()
	return nil
}
func (u *UserPayload) set_updated_at() error {
	u.updated_at = time.Now()
	return nil
}
func (u *UserPayload) set_verification_token() error {
	u.is_verified = false
	u.verification_token = uuid.New()
	u.verification_token_expires_at = time.Now().Add(time.Hour * 24)
	return nil
}
func (u *UserPayload) set_reset_token() error {
	u.reset_token = uuid.New()
	u.reset_token_expires_at = time.Now().Add(time.Hour * 24)
	return nil
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

func (u *User) insert(up UserPayload, d UserPayload) error {
	return nil
}
func (u *User) update(up UserPayload, d UserPayload) error {
	return nil
}
func (u *User) find_by_reset_token(up UserPayload, d UserPayload) error {
	return nil
}
func (u *User) find_by_email(up UserPayload, d UserPayload) error {
	return nil
}
func (u *User) find_by_id(up UserPayload, d UserPayload) error {
	return nil
}
func (u *User) activate(up UserPayload, d UserPayload) error {
	return nil
}
func (u *User) delete(up UserPayload, d UserPayload) error {
	return nil
}
func (u *User) deleteExpired(up UserPayload, d UserPayload) error {
	return nil
}
func (u *User) export(up UserPayload, d UserPayload) error {
	return nil
}
