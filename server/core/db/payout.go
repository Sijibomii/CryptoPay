package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sijibomii/cryptopay/core/models"
)

func insertPayout(conn *gorm.DB, payload models.Payout) models.Payout {
	result := conn.Create(&payload)
	if err := result.Error; err != nil {
		// panic(err)
		fmt.Printf(" errorr %+s\n", result.Error)
	}
	return payload
}
