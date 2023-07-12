package util

import (
	"crypto/rsa"
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

func (payload *JwtPayload) Encode(jwtPrivate *rsa.PrivateKey) (string, error) {

	t := jwt.NewWithClaims(jwt.SigningMethodRS256,
		jwt.MapClaims{
			"client_id":  payload.Client.ID,
			"store_id":   payload.Client.Store_id,
			"expires_at": payload.Expires_at,
		})

	return t.SignedString(jwtPrivate)
}

func DecodeJWT(tokenString string, jwtPrivate *rsa.PrivateKey) (*JwtPayload, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the token's signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtPrivate.Public(), nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Extract the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	client_id := claims["client_id"].(string)
	store_id := claims["store_id"].(string)
	expires_at := claims["expires_at"].(string)

	client_uuid, err := uuid.Parse(client_id)

	store_uuid, err := uuid.Parse(store_id)

	layout := "2006-01-02 15:04:05"

	time, err := time.Parse(layout, expires_at)

	// Map the claims to the JwtPayload struct
	payload := &JwtPayload{
		Client: client.Client{
			ID:       client_uuid,
			Store_id: store_uuid,
		},
		Expires_at: time, //time.Unix(int64(claims["expires_at"].(float64)), 0),
	}

	return payload, nil
}
