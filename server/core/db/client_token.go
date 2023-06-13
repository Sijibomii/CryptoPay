package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sijibomii/cryptopay/core/models"
)

func insertClientToken(conn *gorm.DB, payload models.ClientToken) models.ClientToken {
	result := conn.Create(payload)
	if err := result.Error; err != nil {
		panic(err)
	}
	return payload
}
