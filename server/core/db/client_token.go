package db

import (
	"github.com/google/uuid"
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

func findClientTokensByStore(conn *gorm.DB, store_id uuid.UUID, limit, offset int64) []models.ClientToken {
	var tokens []models.ClientToken

	result := conn.
		Where("store_id = ?", store_id).
		Limit(int(limit)).
		Offset(int(offset)).
		Find(&tokens)

	if result.Error != nil {
		panic(result.Error)
	}

	return tokens
}

func findClientTokenById(conn *gorm.DB, id uuid.UUID) models.ClientToken {
	token := models.ClientToken{}
	if err := conn.Where("id = ?", id).First(&token).Error; err != nil {
		panic(err)
	}

	return token
}

func findClientTokenByTokenAndDomain(conn *gorm.DB, token uuid.UUID, domain string) models.ClientToken {
	res := models.ClientToken{}
	if err := conn.Where("token = ? AND domain = ?", token, domain).First(&token).Error; err != nil {
		panic(err)
	}

	return res
}

func deleteClientToken(conn *gorm.DB, id uuid.UUID) int64 {
	result := conn.
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Delete(&models.ClientToken{})

	if result.Error != nil {
		panic(result.Error)
	}
	return result.RowsAffected
}
