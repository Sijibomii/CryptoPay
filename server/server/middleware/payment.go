package middleware

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sijibomii/cryptopay/server/dao"
	"github.com/sijibomii/cryptopay/server/util"
)

func PaymentMiddleware(appState *util.AppState) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// util.EnableCors(&w)
			auth := r.Header.Get("Authorization")

			parts := strings.Split(auth, " ")

			var token string

			if len(parts) >= 2 && parts[0] == "Bearer" {
				token = parts[1]
			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if token == "" {
				// Handle unauthorized access
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			pattern := `^/p/payments/[^/]+/status$`

			// Create a regex object from the pattern
			regex := regexp.MustCompile(pattern)

			// Check if the URL path matches the regex pattern
			if regex.MatchString(r.URL.Path) || r.URL.Path == "/p/vouchers" {

				// session token is added here instead of api token
				key := appState.PrivateKey

				payload, err := util.DecodeJWT(token, key)

				// check if payment has not expired

				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				ctx := r.Context()
				ctx = context.WithValue(ctx, "Payload", payload)
				r = r.WithContext(ctx)
				//fmt.Printf("PAYMENT STATUS REQUEST 2!! \n")

				// Call the next handler
				next.ServeHTTP(w, r)

			} else if r.URL.Path == "/p/payments" {
				fmt.Print("PAYMENTTTTTTT ############### \n")
				// create payment
				clientToken, tokenErr := uuid.Parse(token)

				if tokenErr != nil {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				originHeader := r.Header.Get("Origin")
				// if originHeader == "" {

				// 	http.Error(w, "invalid origin header", http.StatusUnauthorized)
				// 	return
				// }

				originHeaderParts := strings.Split(originHeader, "://")
				// if len(originHeaderParts) != 2 {
				// 	http.Error(w, "invalid origin header", http.StatusUnauthorized)
				// 	return
				// }

				//fmt.Print("\n")
				// domain := strings.TrimSuffix(originHeaderParts[1], "/")

				// // domain is the store it will be used i.e www.amazon.com
				client_token, err := dao.GetClientTokenByToken(appState.Engine, appState.Postgres, clientToken, "localhost:3000")

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				// Pass the session object to the next handler
				ctx := r.Context()
				ctx = context.WithValue(ctx, "Ctoken", client_token)
				r = r.WithContext(ctx)

				fmt.Print("\n ORIGIN HEADERRR", originHeader)
				fmt.Print("\n ORIGIN HEADERRR PARTSSSS", originHeaderParts)
				// Call the next handler
				next.ServeHTTP(w, r)
			} else {
				//fmt.Print("NOT FOUNDDDD \n")
				http.NotFound(w, r)
			}

		})
	}
}
