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

func findAllConfirmedPayout(conn *gorm.DB, block_number int, cryto string) []models.Payout {
	var payouts []models.Payout
	result := conn.
		Where("status = ?", "pending").
		Where("crypto = ?", cryto).
		Where("block_height_required < ?", block_number).
		Find(&payouts)

	if result.Error != nil {
		panic(result.Error)
	}

	return payouts
}
