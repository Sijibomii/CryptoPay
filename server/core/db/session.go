package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sijibomii/cryptopay/core/models"
)

func insertSession(conn *gorm.DB, payload models.Session) models.Session {
	result := conn.Create(&payload)
	if err := result.Error; err != nil {
		// panic(err)
		//fmt.Printf(" errorr %+s\n", result.Error)
	}
	return payload
}

func getSessionByToken(conn *gorm.DB, token string) models.Session {
	session := models.Session{}

	if err := conn.Where("token = ?", token).First(&session).Error; err != nil {
		//fmt.Printf("################# error: %s\n", err.Error())
		// panic(err)
	}
	// //fmt.Printf("token: %s", token)
	// //fmt.Printf("############ SESSION: %v", session)
	return session
}
