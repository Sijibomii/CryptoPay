package db

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/sijibomii/cryptopay/core/models"
)

func insertPayout(conn *gorm.DB, payload models.Payout) models.Payout {

	result := conn.Create(&payload)
	if err := result.Error; err != nil {
		// panic(err)
		//fmt.Printf(" errorr %+s\n", result.Error)
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

func updatePayoutwithPayment(conn *gorm.DB, payout_id uuid.UUID, payout_payload models.PayoutPayload, payment_payload models.PaymentPayload) models.Payout {
	tx := conn.Begin()

	updatePayment(conn, payout_payload.Payment_id, payment_payload.ToPayment())

	// get payout by id
	if err := tx.Model(models.Payout{}).Where("id = ?", payout_id).Updates(payout_payload.ToPayout()).Error; err != nil {
		tx.Rollback()
		panic("failed to update payout")
	}

	var payout_ models.Payout

	if err := tx.First(&payout_, 1).Error; err != nil {
		tx.Rollback()
		panic("failed to fetch payout")
	}

	// Commit the transaction if both actions succeed
	if err := tx.Commit().Error; err != nil {
		panic("failed to commit transaction")
	}

	return payout_
}
