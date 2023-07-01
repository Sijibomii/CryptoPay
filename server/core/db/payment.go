package db

import (
	"fmt"

	"github.com/google/uuid"
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

func findPaymentById(conn *gorm.DB, id uuid.UUID) models.Payment {
	payment := models.Payment{}
	if err := conn.Where("id = ?", id).First(&payment).Error; err != nil {
		panic(err)
	}

	return payment
}

// not pending
func findAllPayementByAddresses(conn *gorm.DB, addresses []string, cryto string) []models.Payment {
	var payments []models.Payment

	result := conn.
		Where("address IN (?)", addresses).
		Where("crypto = ?", cryto).
		Find(&payments)

	if result.Error != nil {
		panic(result.Error)
	}

	return payments
}

func updatePayment(conn *gorm.DB, id uuid.UUID, payload models.Payment) (models.Payment, error) {
	payment := models.Payment{}
	if err := conn.Where("id = ?", id).First(&payment).Error; err != nil {
		return payload, err
	}

	if err := conn.Save(&payload).Error; err != nil {
		return payload, err
	}

	return payload, nil
}
