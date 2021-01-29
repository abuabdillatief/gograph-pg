package middlewares

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/abuabdillatief/gograph-tutorial/database"
	"github.com/abuabdillatief/gograph-tutorial/graph/model"
	"github.com/dgrijalva/jwt-go"
	jwtRequest "github.com/dgrijalva/jwt-go/request"
	"github.com/pkg/errors"
)

//CurrentUserType ...
type CurrentUserType string

//CurrentUser ...
const CurrentUser = CurrentUserType("currentUser")

var authHeaderExtractor = &jwtRequest.PostExtractionFilter{
	Extractor: jwtRequest.HeaderExtractor{"Authorization"},
	Filter:    bearerPrefixFromToken,
}

var authExtractor = &jwtRequest.MultiExtractor{
	authHeaderExtractor,
	jwtRequest.ArgumentExtractor{"access_token"},
}

//AuthMiddleware ...
func AuthMiddleware(DB database.UsersRepo) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := parseToken(r)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			//here means, we want to create a claim
			//specifically, a map claim
			//in jwt doc, there are 2 types of claims
			//standard claims and map claims
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				next.ServeHTTP(w, r)
				return
			}
			user, err := DB.GetUserByID(claims["jti"].(string))
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			ctx := context.WithValue(r.Context(), CurrentUser, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

//GetCurrentUserFromContext ...
func GetCurrentUserFromContext(ctx context.Context) (*model.User, error) {
	noUserError := errors.New("There's no user in context")
	if ctx.Value(CurrentUser) == nil {
		return nil, noUserError
	}
	user, ok := ctx.Value(CurrentUser).(*model.User)
	if !ok || user.ID == "" {
		return nil, noUserError
	}
	return user, nil
}

func parseToken(r *http.Request) (*jwt.Token, error) {
	parsed, err := jwtRequest.ParseFromRequest(r, authExtractor, func(token *jwt.Token) (i interface{}, err error) {
		t := []byte(os.Getenv("JWT_SECRET"))
		return t, nil
	})
	if err != nil {
		return nil, err
	}
	return parsed, errors.Wrap(err, "parseToken error: ")
}

func bearerPrefixFromToken(token string) (string, error) {
	bearer := "BEARER"
	if len(token) > len(bearer) && strings.ToUpper(token[0:len(bearer)]) == bearer {
		return token[len(bearer)+1:], nil
	}
	return token, nil
}
