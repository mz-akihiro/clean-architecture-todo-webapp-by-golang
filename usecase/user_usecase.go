package usecase

import (
	"clean-architecture-todo-webapp-by-golang/model"
	"clean-architecture-todo-webapp-by-golang/repository"
	"database/sql"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	LogIn(user model.User) (int, string, error)
	SignUp(user model.User) (int, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) LogIn(user model.User) (int, string, error) {
	var (
		storedPassword string
		userId         int
	)
	if user.Email == "" || user.Password == "" {
		return http.StatusBadRequest, "", errors.New("invalid JSON format")
	}
	if err := uu.ur.GetUserByEmail(&storedPassword, &userId, user.Email); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusUnauthorized, "", err // セキュリティの観点から403は使わない
		} else {
			return http.StatusInternalServerError, "", err
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password)); err != nil {
		return http.StatusUnauthorized, "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,                               // 後々taskを追加・削除する際などに使う用に、ペイロードにuserIdを入れる
		"exp":     time.Now().Add(time.Hour * 6).Unix(), // jwtトークンの有効期限(6時間)
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	return http.StatusOK, tokenString, nil
}

func (uu *userUsecase) SignUp(user model.User) (int, error) {
	var count int
	if user.Email == "" || user.Password == "" {
		return http.StatusBadRequest, errors.New("invalid JSON format")
	}

	// 既に登録済みアカウントかどうか確認
	if err := uu.ur.SearchUserByEmail(&count, user.Email); err != nil {
		return http.StatusInternalServerError, err
	}
	if count > 0 {
		return http.StatusConflict, errors.New("alredy sign up")
	}

	// 未登録であればpasswordをhash化し、登録する
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if err := uu.ur.CreateUser(user.Email, string(hash)); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil

}
