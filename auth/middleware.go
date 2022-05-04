package auth

import (
	"context"
	"net/http"

	"github.com/DavidHODs/TechHUB-goGraph/graph/model"
	myDb "github.com/DavidHODs/TechHUB-goGraph/postgres"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	email string
}

func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			tokenStr := header
			email, err := ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "invalid token", http.StatusForbidden)
				return
			}

			user := model.User{}
			id, _, name, err := myDb.GetUserDetailsByEmail(email)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.ID = id
			user.Name = name
			user.Email = email

			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value("AuthToken").(*model.User)
	return raw
}