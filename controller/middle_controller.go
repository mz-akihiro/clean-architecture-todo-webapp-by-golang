package controller

import (
	"clean-architecture-todo-webapp-by-golang/model"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

type IMiddleController interface {
	AuthMiddleware(next http.HandlerFunc) http.HandlerFunc
}

type middleController struct {
}

func NewMiddleController() IMiddleController {
	return &middleController{}
}

func (mc *middleController) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		value, err := r.Cookie("token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Request does not contain a token")
			return
		}

		// Cookie内のtokenを検証する
		token, err := jwt.Parse(value.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Unauthorized token")
			return
		} else {
			userID := int(token.Claims.(jwt.MapClaims)["user_id"].(float64))
			ctx := context.WithValue(r.Context(), model.ContextKey{UserId: "user_id"}, userID)
			defer next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}
