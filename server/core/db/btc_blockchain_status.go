package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sijibomii/cryptopay/core/models"
)

func insertBtcBlockChainStatus(conn *gorm.DB, payload models.BtcBlockChainStatus) models.BtcBlockChainStatus {
	result := conn.Create(&payload)
	if err := result.Error; err != nil {
		panic(err)
		fmt.Printf(" errorr %+s\n", result.Error)
	}
	return payload
}

func findBtcBlockChainStatus(conn *gorm.DB, network string) models.BtcBlockChainStatus {

	btcs := models.BtcBlockChainStatus{}

	if err := conn.Where("network = ?", network).
		Order("created_at DESC").
		First(&btcs).Error; err != nil {
		panic(err)
	}

	return btcs
}
