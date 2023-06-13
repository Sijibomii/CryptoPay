package db

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/sijibomii/cryptopay/core/models"
)

func insertStore(conn *gorm.DB, payload models.Store) (models.Store, error) {
	result := conn.Create(payload)
	if err := result.Error; err != nil {
		return payload, err
	}
	return payload, nil
}

func updateStore(conn *gorm.DB, id uuid.UUID, payload models.Store) (models.Store, error) {
	store := models.Store{}
	if err := conn.Where("id = ? AND deleted_at IS NULL", id).First(&store).Error; err != nil {
		return payload, err
	}

	if err := conn.Save(&payload).Error; err != nil {
		return payload, err
	}

	return payload, nil
}

func findStoreById(conn *gorm.DB, id uuid.UUID) (models.Store, error) {
	store := models.Store{}
	if err := conn.Where("id = ? AND deleted_at IS NULL", id).First(&store).Error; err != nil {
		return store, err
	}

	return store, nil
}

// this includes stores that have been soft deleted
func findStoreByIdWithDeleted(conn *gorm.DB, id uuid.UUID) (models.Store, error) {
	store := models.Store{}
	if err := conn.Where("id = ?", id).First(&store, id).Error; err != nil {
		return store, err
	}

	return store, nil
}

func findByOwner(conn *gorm.DB, ownerID uuid.UUID, limit, offset int64) ([]models.Store, error) {
	var stores []models.Store

	result := conn.
		Where("owner_id = ?", ownerID).
		Where("deleted_at IS NULL").
		Limit(int(limit)).
		Offset(int(offset)).
		Find(&stores)

	if result.Error != nil {
		return nil, result.Error
	}

	return stores, nil
}

func delete(conn *gorm.DB, id uuid.UUID) (int64, error) {
	result := conn.
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Delete(&models.Store{})

	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func SoftDelete(conn *gorm.DB, id uuid.UUID) error {
	payload := models.StorePayload{}
	payload.Set_deleted_at()

	result := conn.
		Model(&models.Store{}).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Updates(payload)

	if result.Error != nil {
		return result.Error
	}

	// delete client tokens by sotr id

	return nil
}

func SoftDeleteByOwnerID(conn *gorm.DB, ownerID uuid.UUID) error {
	payload := models.StorePayload{}
	payload.Set_deleted_at()

	result := conn.
		Model(&models.Store{}).
		Where("owner_id = ?", ownerID).
		Where("deleted_at IS NULL").
		Updates(payload)

	if result.Error != nil {
		return result.Error
	}

	var deletedStores []models.Store
	if err := conn.
		Where("owner_id = ?", ownerID).
		Find(&deletedStores).Error; err != nil {
		return err
	}

	// for _, store := range deletedStores {
	// 	// delete client tokens of each store
	// }

	return nil
}
