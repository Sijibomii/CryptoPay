package models

import (
	"errors"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/google/uuid"
)

type Session struct {
	ID             string                 `json:"id"`
	Token          string                 `json:"token"`
	UserID         uuid.UUID              `json:"user_id"`
	Props          map[string]interface{} `gorm:"type:jsonb"`
	CreatedAt      time.Time              `json:"create_at,omitempty"`
	UpdatedAt      time.Time              `json:"update_at,omitempty"`
	LastActivityAt time.Time              `json:"last_activity_at"`
}

func (sp *Session) Set_created_at() error {
	sp.CreatedAt = time.Now()
	return nil
}

func (sp *Session) Set_updated_at() error {
	sp.UpdatedAt = time.Now()
	return nil
}
func (sp *Session) Set_LastActivity_at() error {
	sp.LastActivityAt = time.Now()
	return nil
}

type InsertSessionMessage struct {
	Payload Session
}

type UpdateSessionMessage struct {
	Payload Session
	Id      uuid.UUID
}

type GetSessionByTokenMessage struct {
	Token string
}

func InsertSession(e *actor.Engine, conn *actor.PID, d *Session) (Session, error) {
	d.Set_created_at()
	var resp = e.Request(conn, InsertSessionMessage{
		Payload: *d,
	}, time.Millisecond*100)
	res, err := resp.Result()
	if err != nil {
		return Session{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(Session)

	if !ok {
		return Session{}, errors.New("An error occured!")
	}

	return myStruct, nil
}

func GetSessionByToken(e *actor.Engine, conn *actor.PID, token string) (Session, error) {

	var resp = e.Request(conn, GetSessionByTokenMessage{
		Token: token,
	}, time.Millisecond*100)

	res, err := resp.Result()

	if err != nil {
		return Session{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(Session)

	if !ok {
		return Session{}, errors.New("An error occured!")
	}

	return myStruct, nil
}

// func UpdateSession(e *actor.Engine, conn *actor.PID, id uuid.UUID, d UserPayload) (User, error) {
// 	var resp = e.Request(conn, UpdateUserMessage{
// 		Payload: d.ToUser(),
// 		Id:      id,
// 	}, 500)

// 	res, err := resp.Result()
// 	if err != nil {
// 		return User{}, errors.New("An error occured!")
// 	}
// 	myStruct, ok := res.(User)

// 	if !ok {
// 		return User{}, errors.New("An error occured!")
// 	}

// 	return myStruct, nil

// }
