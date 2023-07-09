package middleware

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sijibomii/cryptopay/server/dao"
	"github.com/sijibomii/cryptopay/server/util"
)

func AuthMiddleware(appState *util.AppState) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// util.EnableCors(&w)
			token := r.Header.Get("Authorization")
			if token == "" {
				// Handle unauthorized access
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			session, err := dao.GetSessionByToken(appState.Engine, appState.Postgres, token)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			user, userErr := dao.GetUserById(appState.Engine, appState.Postgres, session.UserID)

			if userErr != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Pass the session object to the next handler
			ctx := r.Context()
			ctx = context.WithValue(ctx, "user", user)
			r = r.WithContext(ctx)

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}
