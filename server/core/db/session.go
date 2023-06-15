package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sijibomii/cryptopay/core/models"
)

func insertSession(conn *gorm.DB, payload models.Session) models.Session {
	result := conn.Create(payload)
	if err := result.Error; err != nil {
		panic(err)
	}
	return payload
}
