package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sijibomii/cryptopay/server/dao"
	"github.com/sijibomii/cryptopay/server/util"
)

func PaymentMiddleware(appState *util.AppState) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
			clientToken, tokenErr := uuid.Parse(token)

			if tokenErr != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			originHeader := r.Header.Get("Origin")
			if originHeader == "" {
				http.Error(w, "invalid origin header", http.StatusUnauthorized)
				return
			}

			originHeaderParts := strings.Split(originHeader, "://")
			if len(originHeaderParts) != 2 {
				http.Error(w, "invalid origin header", http.StatusUnauthorized)
				return
			}

			domain := strings.TrimSuffix(originHeaderParts[1], "/")

			// // domain is the store it will be used i.e www.amazon.com
			client_token, err := dao.GetClientTokenByToken(appState.Engine, appState.Postgres, clientToken, domain)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Pass the session object to the next handler
			ctx := r.Context()
			ctx = context.WithValue(ctx, "Ctoken", client_token)
			r = r.WithContext(ctx)

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}
