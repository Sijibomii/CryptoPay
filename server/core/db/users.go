package db

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/sijibomii/cryptopay/core/models"
)

func insertUser(conn *gorm.DB, payload models.User) (models.User, error) {
	result := conn.Create(payload)
	// created_at?
	if err := result.Error; err != nil {
		return payload, err
	}
	return payload, nil
}

func updateUser(conn *gorm.DB, id uuid.UUID, payload models.User) (models.User, error) {
	user := models.User{}
	if err := conn.Where("id = ?", id).First(&user).Error; err != nil {
		return payload, err
	}

	if err := conn.Save(&payload).Error; err != nil {
		return payload, err
	}

	return payload, nil
}

func findUserByEmail(conn *gorm.DB, email string) (models.User, error) {
	var users []models.User
	if err := conn.Where("email = ?", email).Find(&users).Error; err != nil {
		return users[0], err
	}

	return users[0], nil
}

func findUserById(conn *gorm.DB, id uuid.UUID) models.User {
	user := models.User{}
	if err := conn.Where("id = ?", id).First(&user).Error; err != nil {
		panic(err)
	}

	return user
}

func findUserByResetToken(conn *gorm.DB, token uuid.UUID) models.User {
	var users []models.User
	if err := conn.Where("reset_token = ?", token).Find(&users).Error; err != nil {
		panic(err)
	}

	return users[0]
}

func activateUser(conn *gorm.DB, token uuid.UUID) models.User {
	var user models.User
	now := time.Now()

	if err := conn.Where("verification_token = ? AND is_activated = ? AND verification_token_expires_at > ?", token, false, now).First(&user).Error; err != nil {
		panic(err)
	}

	user.Is_verified = true

	result := conn.Update(user)

	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected == 0 {
		panic(gorm.ErrRecordNotFound)
	}
	return user
}

func deleteUser(conn *gorm.DB, id uuid.UUID) models.User {
	result := conn.Delete(&models.User{}, id)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected == 0 {
		panic(fmt.Errorf("no record found with ID %d", id))
	}
	return models.User{}
}

func deleteExpiredUser(conn *gorm.DB, email string) models.User {
	var user models.User
	now := time.Now()

	if err := conn.Where("email = ? AND is_activated = ? AND verification_token_expires_at < ?", email, false, now).First(&user).Error; err != nil {
		panic(err)
	}

	result := conn.Delete(user)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected == 0 {
		panic(gorm.ErrRecordNotFound)
	}
	return user
}
