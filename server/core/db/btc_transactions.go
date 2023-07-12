package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sijibomii/cryptopay/core/models"
)

func insertBtcTransaction(conn *gorm.DB, payload models.BtcTransaction) models.BtcTransaction {
	result := conn.Create(&payload)
	if err := result.Error; err != nil {
		// panic(err)
		//fmt.Printf(" errorr %+s\n", result.Error)
	}
	return payload
}
