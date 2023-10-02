package controller

import (
	"clean-architecture-todo-webapp-by-golang/model"
	"clean-architecture-todo-webapp-by-golang/usecase"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type IUserController interface {
	LogIn(w http.ResponseWriter, r *http.Request)
	LogOut(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	uu usecase.IUserUsecase
}

/* type userController2 struct {
	uu usecase.IUserUsecase
} */

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) LogIn(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid JSON format")
		return
	}
	statusCode, tokenString, err := uc.uu.LogIn(user)
	if err != nil {
		w.WriteHeader(statusCode)
		fmt.Fprintln(w, "Failed to create token")
		return
	}
	cookie := &http.Cookie{}
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(12 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	http.SetCookie(w, cookie)
}

func (uc *userController) LogOut(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{}

	cookie.Name = "token"
	cookie.Value = ""           // valueを上書きする
	cookie.Expires = time.Now() // 現在時刻が期限なのですぐにクッキーから削除される
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode

	http.SetCookie(w, cookie)
	fmt.Fprintln(w, "Delete cookie")
}

func (uc *userController) SignUp(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid JSON format")
		return
	}
	if statusCode, err := uc.uu.SignUp(user); err != nil {
		w.WriteHeader(statusCode)
		fmt.Fprintln(w, err)
		return
	}
}
