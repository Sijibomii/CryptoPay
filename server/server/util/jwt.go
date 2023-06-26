package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sijibomii/cryptopay/server/client"
)

type JwtPayload struct {
	Client     client.Client
	Expires_at time.Time
}

func (payload *JwtPayload) Encode(jwtPrivate string) (string, error) {
	key := "" /* Load key from somewhere, for example a file */
	t := jwt.NewWithClaims(jwt.SigningMethodRS256,
		jwt.MapClaims{
			"client": jwt.MapClaims{
				"id":       payload.Client.ID,
				"store_id": payload.Client.Store_id,
			},
			"expires_at": payload.Expires_at,
		})
	return t.SignedString(key)
}
