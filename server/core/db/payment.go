package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sijibomii/cryptopay/core/models"
)

func insertPayment(conn *gorm.DB, payload models.Payment) models.Payment {
	result := conn.Create(&payload)
	if err := result.Error; err != nil {
		// panic(err)
		fmt.Printf(" errorr %+s\n", result.Error)
	}
	return payload
}
