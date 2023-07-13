package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sijibomii/cryptopay/core/models"
)

func insertBtcBlockChainStatus(conn *gorm.DB, payload models.BtcBlockChainStatus) models.BtcBlockChainStatus {
	result := conn.Create(&payload)
	if err := result.Error; err != nil {
		fmt.Printf(" errorr %+s\n", result.Error)
		panic(err)
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

func updateBtcBlockChainStatusByNetwork(conn *gorm.DB, network string, block_height int) models.BtcBlockChainStatus {

	var btcStatus models.BtcBlockChainStatus
	if err := conn.Where("network = ?", network).First(&btcStatus).Error; err != nil {
		fmt.Printf("\n err: %s \n", err.Error())
		panic("error getting network before update")
	}

	// Update the height field
	btcStatus.Block_Height = block_height
	if err := conn.Save(&btcStatus).Error; err != nil {
		fmt.Printf("\n err: %s \n", err.Error())
		panic("unable to save btc network status")
	}
	return btcStatus
}
