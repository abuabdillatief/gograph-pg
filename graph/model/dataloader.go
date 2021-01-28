package model

import (
	"context"
	"net/http"
	"time"

	"github.com/go-pg/pg/v9"
)

type contextKey string

const userLoaderKey = contextKey("userloader")

//DataloaderMiddlewareDB ...
//it takes http.Handler because it is a middleware
//and in the end we want to hand it over the next handler
func DataloaderMiddlewareDB(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userLoaderValue := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*User, []error) {
				var users []*User
				err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()
				if err != nil {
					return nil, []error{err}
				}

				u := make(map[string]*User)
				for _, user := range users {
					u[user.ID] = user
				}
				for i, id := range ids {
					users[i] = u[id]
				}

				return users, []error{err}
			},
		}
		ctx := context.WithValue(r.Context(), userLoaderKey, &userLoaderValue)
		//http.Handler is an interface that owns a method called ServeHTT
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//GetUserLoader ...
func GetUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userLoaderKey).(*UserLoader)
}
