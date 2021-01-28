package model

import (
	"context"
	"net/http"
	"time"

	"github.com/go-pg/pg/v9"
)

const userLoaderKey = "userloader"

func DataloaderMiddlewareDB(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userLoader := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*User, []error) {
				var users []*User
				err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()
				if err != nil {
					return nil, []error{err}
				}
				return users, []error{err}
			},
		}

		ctx := context.WithValue(r.Context(), userLoaderKey, &userLoader)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userLoaderKey).(*UserLoader)
}
