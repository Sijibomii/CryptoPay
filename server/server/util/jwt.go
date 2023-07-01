package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sijibomii/cryptopay/server/client"
)

type JwtPayload struct {
	Client     client.Client
	Expires_at time.Time
}

func (payload *JwtPayload) Encode(jwtPrivate string) (string, error) {

	t := jwt.NewWithClaims(jwt.SigningMethodRS256,
		jwt.MapClaims{
			"client_id":  payload.Client.ID,
			"store_id":   payload.Client.Store_id,
			"expires_at": payload.Expires_at,
		})
	return t.SignedString(jwtPrivate)
}

func DecodeJWT(tokenString, secret_key string) (*JwtPayload, error) {

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret_key), nil
	})

	if err != nil {
		fmt.Printf("error decoding jwt token %s \n", tokenString)
	}

	client_id := claims["client_id"].(string)
	store_id := claims["store_id"].(string)
	expires_at := claims["expires_at"].(string)

	client_uuid, err := uuid.Parse(client_id)

	store_uuid, err := uuid.Parse(store_id)

	layout := "2006-01-02 15:04:05"

	parsedTime, err := time.Parse(layout, expires_at)

	return &JwtPayload{
		Client: client.Client{
			ID:       client_uuid,
			Store_id: store_uuid,
		},
		Expires_at: parsedTime,
	}, nil
}
