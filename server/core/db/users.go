package db

import (
	"fmt"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/sijibomii/cryptopay/core/models"
)

func insert(conn *gorm.DB, payload models.User) (models.User, error) {
	result := conn.Create(payload)
	if err := result.Error; err != nil {
		return payload, err
	}
	return payload, nil
}

func update(conn *gorm.DB, id uuid.UUID, payload models.User) (models.User, error) {
	user := models.User{}
	if err := conn.First(&user, id).Error; err != nil {
		return payload, err
	}

	if err := conn.Save(&payload).Error; err != nil {
		return payload, err
	}

	return payload, nil
}

func find_by_email(conn *gorm.DB, email string) (models.User, error) {
	var users []models.User
	if err := conn.Where("email = ?", email).Find(&users).Error; err != nil {
		return users[0], err
	}

	return users[0], nil
}

func find_by_id(conn *gorm.DB, id uuid.UUID) (models.User, error) {
	user := models.User{}
	if err := conn.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func find_by_reset_token(conn *gorm.DB, token uuid.UUID) (models.User, error) {
	var users []models.User
	if err := conn.Where("reset_token = ?", token).Find(&users).Error; err != nil {
		return users[0], err
	}

	return users[0], nil
}

func activate(conn *gorm.DB, token uuid.UUID) (models.User, error) {
	var user models.User
	now := time.Now()

	if err := conn.Where("verification_token = ? AND is_activated = ? AND verification_token_expires_at > ?", token, false, now).First(&user).Error; err != nil {
		return user, err
	}

	user.Is_verified = true

	result := conn.Update(user)

	if result.Error != nil {
		return models.User{}, result.Error
	}
	if result.RowsAffected == 0 {
		return models.User{}, gorm.ErrRecordNotFound
	}
	return user, nil
}

func delete(conn *gorm.DB, id uuid.UUID) (models.User, error) {
	result := conn.Delete(&models.User{}, id)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	if result.RowsAffected == 0 {
		return models.User{}, fmt.Errorf("no record found with ID %d", id)
	}
	return models.User{}, nil
}

func delete_expired(conn *gorm.DB, email string) (models.User, error) {
	var user models.User
	now := time.Now()

	if err := conn.Where("email = ? AND is_activated = ? AND verification_token_expires_at < ?", email, false, now).First(&user).Error; err != nil {
		return user, err
	}

	result := conn.Delete(user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	if result.RowsAffected == 0 {
		return models.User{}, gorm.ErrRecordNotFound
	}
	return user, nil
}

type client struct {
	Conn      *gorm.DB
	ServerPID *actor.PID
}

func (c *client) Receive(ctx *actor.Context) {
	switch l := ctx.Message().(type) {
	case actor.Started:
		fmt.Println("User db actor started")

	case models.Insert:
		payload, err := insert(c.Conn, l.Payload)
		if err != nil {
			ctx.Respond(nil)
		}
		ctx.Respond(payload)

	case models.Update:
		update(c.Conn, l.Id, l.Payload)

	case models.FindByEmail:
		find_by_email(c.Conn, l.Email)

	case models.FindById:
		find_by_id(c.Conn, l.Id)

	case models.FindByResetToken:
		find_by_reset_token(c.Conn, l.Token)

	case models.Activate:
		activate(c.Conn, l.Token)

	case models.Delete:
		delete(c.Conn, l.Token)

	case models.DeleteExpired:
		delete_expired(c.Conn, l.Email)

	default:
		fmt.Println("UNKNOWN MESSAGE TO USER DB")
	}
}
