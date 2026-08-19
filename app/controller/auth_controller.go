package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"api-backend/sample/app/application"
	"api-backend/sample/app/application/env"
	"api-backend/sample/app/controller/dto"
	"api-backend/sample/app/controller/middleware"
	"api-backend/sample/app/models"
	"api-backend/sample/app/service/auth"
	app_http "api-backend/sample/app/service/http"

	"github.com/golang-jwt/jwt/v5"
)

type AuthResponse struct {
	UserId       uint   `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func Register(rw http.ResponseWriter, r *http.Request) {
	var registerRequest dto.AuthUser

	err := json.NewDecoder(r.Body).Decode(&registerRequest)

	if err != nil {
		fmt.Fprint(rw, "error occured")
		fmt.Fprint(rw, err.Error())
	}

	existingUser := models.User{}

	var userExists bool
	existsQuery := fmt.Sprintf(
		"SELECT EXISTS(SELECT name FROM %s WHERE name = @name) AS user_exists",
		models.UserTable,
	)
	application.DB.Raw(
		existsQuery,
		sql.Named("name", registerRequest.Name),
	).Scan(&userExists)

	application.DB.Where("name = ?", registerRequest.Name).First(&existingUser)

	if userExists {
		app_http.SetJsonResponse(
			rw,
			dto.ErrorResponse{
				Error: "Incorrect details",
			},
		)

		return
	}

	accessTokenExpiryTimestamp := time.Now().Add(time.Hour).Unix()
	refreshTokenExpiryTimestamp := time.Now().Add(time.Hour * 24).Unix()

	hashedPassword := auth.HashPassword([]byte(registerRequest.Password))
	accessToken, accessTokenError := auth.GenerateJWT(
		registerRequest.Name,
		accessTokenExpiryTimestamp,
		[]byte(env.JwtSecret.GetValue()),
	)
	refreshToken, refreshTokenError := auth.GenerateJWT(
		registerRequest.Name,
		refreshTokenExpiryTimestamp,
		[]byte(env.JwtSecret.GetValue()),
	)

	if accessTokenError != nil || refreshTokenError != nil {
		// todo: logging
		fmt.Fprint(rw, "something wrong happened", err.Error())
	}

	user := models.User{Name: registerRequest.Name, Password: hashedPassword, AccessToken: accessToken, RefreshToken: refreshToken}

	application.DB.Create(&user)

	app_http.SetJsonResponse(
		rw,
		AuthResponse{
			UserId:       user.ID,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	)
}

func Login(rw http.ResponseWriter, r *http.Request) {
	var loginRequest dto.AuthUser
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		fmt.Fprint(rw, "error occured")
		fmt.Fprint(rw, err.Error())
	}

	var user models.User
	userErr := application.DB.Where("name = ?", loginRequest.Name).First(&user).Error
	if userErr != nil {
		app_http.SetJsonResponse(
			rw,
			dto.ErrorResponse{
				Error: "Incorrect login details",
			},
		)

		return
	}

	isPasswordValid := auth.VerifyPassword([]byte(loginRequest.Password), []byte(user.Password))
	if isPasswordValid == false {
		app_http.SetJsonResponse(
			rw,
			dto.ErrorResponse{
				Error: "Incorrect login details",
			},
		)

		return
	}

	accessTokenExpiryTimestamp := time.Now().Add(time.Hour).Unix()
	refreshTokenExpiryTimestamp := time.Now().Add(time.Hour * 24).Unix()

	accessToken, accessTokenError := auth.GenerateJWT(
		loginRequest.Name,
		accessTokenExpiryTimestamp,
		[]byte(env.JwtSecret.GetValue()),
	)
	if accessTokenError != nil {
		app_http.SetJsonResponse(
			rw,
			dto.ErrorResponse{
				Error: "Auth error",
			},
		)

		return
	}

	refreshToken, refreshTokenError := auth.GenerateJWT(
		loginRequest.Name,
		refreshTokenExpiryTimestamp,
		[]byte(env.JwtSecret.GetValue()),
	)
	if refreshTokenError != nil {
		app_http.SetJsonResponse(
			rw,
			dto.ErrorResponse{
				Error: "Auth error",
			},
		)

		return
	}

	application.DB.Model(&user).Updates(models.User{AccessToken: accessToken, RefreshToken: refreshToken})

	app_http.SetJsonResponse(
		rw,
		AuthResponse{
			UserId:       user.ID,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	)
}

func Logout(rw http.ResponseWriter, r *http.Request) {
	// token := r.Context().Value(middleware.JwtTokenContextKey).(*jwt.Token)
	// fmt.Printf("context value: %+v\n", token)
	fmt.Fprintln(rw, "logging out...")
}

func RefreshToken(rw http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(middleware.JwtTokenContextKey).(*jwt.Token)
	fmt.Printf("context value: %+v\n", token)
	fmt.Fprintln(rw, "logging out...")
}
