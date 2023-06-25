package db

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/sijibomii/cryptopay/core/models"
)

func insertStore(conn *gorm.DB, payload models.Store) models.Store {

	result := conn.Create(&payload)
	if err := result.Error; err != nil {
		panic(err)
	}
	return payload
}

func updateStore(conn *gorm.DB, id uuid.UUID, payload models.Store) models.Store {
	store := models.Store{
		ID: payload.ID,
	}

	fieldsToUpdate := map[string]interface{}{}

	if payload.Name != "" {
		fieldsToUpdate["Name"] = payload.Name
	}

	if payload.Description != "" {
		fieldsToUpdate["Description"] = payload.Description
	}

	result := conn.Model(&store).Updates(fieldsToUpdate)

	if err := result.Error; err != nil {
		panic(err)
	}
	return payload
}

func findStoreById(conn *gorm.DB, id uuid.UUID) models.Store {

	store := models.Store{}
	if err := conn.Where("id = ? AND NOT del", id).First(&store).Error; err != nil {
		panic(err)
	}

	return store
}

// this includes stores that have been soft deleted
func findStoreByIdWithDeleted(conn *gorm.DB, id uuid.UUID) models.Store {
	store := models.Store{}
	if err := conn.Where("id = ?", id).First(&store, id).Error; err != nil {
		panic(err)
	}

	return store
}

func findStoreByOwner(conn *gorm.DB, ownerID uuid.UUID, limit, offset int64) []models.Store {
	var stores []models.Store

	result := conn.
		Where("owner_id = ?", ownerID).
		Where("del = ?", false).
		Limit(int(limit)).
		Offset(int(offset)).
		Find(&stores)

	if result.Error != nil {
		panic(result.Error)
	}

	return stores
}

func deleteStore(conn *gorm.DB, id uuid.UUID) int64 {
	result := conn.
		Where("id = ?", id).
		Where("NOT del").
		Delete(&models.Store{})

	if result.Error != nil {
		panic(result.Error)
	}
	return result.RowsAffected
}

func softDeleteStore(conn *gorm.DB, id uuid.UUID) bool {
	payload := models.StorePayload{}
	payload.Set_deleted_at()

	result := conn.
		Model(&models.Store{}).
		Where("id = ?", id).
		Where("NOT del").
		Updates(payload)

	if result.Error != nil {
		panic(result.Error)
	}

	// delete client tokens by sotr id

	return true
}

func softDeleteStoreByOwnerID(conn *gorm.DB, ownerID uuid.UUID) bool {
	payload := models.StorePayload{}
	payload.Set_deleted_at()

	result := conn.
		Model(&models.Store{}).
		Where("owner_id = ?", ownerID).
		Where("NOT del").
		Updates(payload)

	if result.Error != nil {
		panic(result.Error)
	}

	var deletedStores []models.Store
	if err := conn.
		Where("owner_id = ?", ownerID).
		Find(&deletedStores).Error; err != nil {
		panic(err)
	}

	// for _, store := range deletedStores {
	// 	// delete client tokens of each store
	// }

	return true
}
