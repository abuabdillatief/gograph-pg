package model

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-pg/pg/v9"
)

const userLoaderKey = "userloader"

//DataloaderMiddlewareDB ...
//it takes http.Handler because it is a middleware
//and in the end we want to hand it over the next handler
func DataloaderMiddlewareDB(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userLoaderValue := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*User, []error) {
				fmt.Println(ids, "<<< ids")
				var users []*User
				err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()
				fmt.Println(users, "<<<< users")
				if err != nil {
					return nil, []error{err}
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
