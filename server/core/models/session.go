package models

import (
	"github.com/google/uuid"
)

type Session struct {
	ID             string                 `json:"id"`
	Token          string                 `json:"token"`
	UserID         string                 `json:"user_id"`
	Props          map[string]interface{} `json:"props"`
	CreatedAt      int64                  `json:"create_at,omitempty"`
	UpdatedAt      int64                  `json:"update_at,omitempty"`
	LastActivityAt int64                  `json:"last_activity_at"`
}

type InsertSessionMessage struct {
	Payload Session
}

type UpdateSessionMessage struct {
	Payload Session
	Id      uuid.UUID
}
