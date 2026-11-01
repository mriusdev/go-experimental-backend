package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"api-backend/sample/app/application/env"
	httperrors "api-backend/sample/app/application/http_errors"
	logserrors "api-backend/sample/app/application/logs_errors"
	"api-backend/sample/app/application/utils"
	"api-backend/sample/app/controller/dto"
	"api-backend/sample/app/controller/middleware"
	"api-backend/sample/app/models"
	"api-backend/sample/app/service/auth"
	app_http "api-backend/sample/app/service/http"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthResponse struct {
	UserId       uint   `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthController struct {
	Logger *slog.Logger
	DB     *gorm.DB
}

func NewAuthController(logger *slog.Logger, db *gorm.DB) *AuthController {
	return &AuthController{
		Logger: logger,
		DB:     db,
	}
}

func (a *AuthController) Register(rw http.ResponseWriter, r *http.Request) {
	var registerRequest dto.AuthUser

	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		app_http.SetJsonResponse(
			rw,
			dto.ErrorResponse{
				Error: httperrors.API_INVALID_REQUEST_JSON,
			},
		)

		return
	}

	existingUser := models.User{}

	var userExists bool
	existsQuery := fmt.Sprintf(
		"SELECT EXISTS(SELECT name FROM %s WHERE name = @name) AS user_exists",
		models.UserTable,
	)
	a.DB.Raw(
		existsQuery,
		sql.Named("name", registerRequest.Name),
	).Scan(&userExists)

	a.DB.Where("name = ?", registerRequest.Name).First(&existingUser)

	if userExists {
		app_http.SetJsonResponse(
			rw,
			dto.ErrorResponse{
				Error: httperrors.AUTH_INVALID_DETAILS,
			},
		)

		return
	}

	hashedPassword := auth.HashPassword([]byte(registerRequest.Password))

	user := models.User{Name: registerRequest.Name, Password: hashedPassword}

	a.DB.Create(&user)

	accessTokenExpiryTimestamp := time.Now().Add(time.Hour).Unix()
	refreshTokenExpiryTimestamp := time.Now().Add(time.Hour * 24).Unix()

	accessToken, accessTokenError := auth.GenerateJWT(
		strconv.FormatUint(uint64(user.ID), utils.DecimalDisplay),
		accessTokenExpiryTimestamp,
		[]byte(env.JwtSecret.GetValue()),
	)
	refreshToken, refreshTokenError := auth.GenerateJWT(
		strconv.FormatUint(uint64(user.ID), utils.DecimalDisplay),
		refreshTokenExpiryTimestamp,
		[]byte(env.JwtSecret.GetValue()),
	)

	if accessTokenError != nil {
		a.Logger.Error(logserrors.AUTH_CONTROLLER_ERROR, "Access token generation error", accessTokenError)
	}
	if refreshTokenError != nil {
		a.Logger.Error(logserrors.AUTH_CONTROLLER_ERROR, "Refresh token generation error", refreshTokenError)
	}

	if accessTokenError != nil || refreshTokenError != nil {
		app_http.SetJsonResponse(
			rw,
			dto.ErrorResponse{
				Error: httperrors.API_GENERIC,
			},
		)

		return
	}

	a.DB.Model(&user).UpdateColumns(models.User{AccessToken: accessToken, RefreshToken: refreshToken})

	app_http.SetJsonResponse(
		rw,
		AuthResponse{
			UserId:       user.ID,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	)
}

func (a *AuthController) Login(rw http.ResponseWriter, r *http.Request) {
	var loginRequest dto.AuthUser
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		app_http.SetJsonResponse(
			rw,
			dto.ErrorResponse{
				Error: httperrors.API_INVALID_REQUEST_JSON,
			},
		)

		return
	}

	var user models.User
	userErr := a.DB.Where("name = ?", loginRequest.Name).First(&user).Error
	if userErr != nil {
		app_http.SetJsonResponse(
			rw,
			dto.ErrorResponse{
				Error: httperrors.AUTH_INVALID_DETAILS,
			},
		)

		return
	}

	isPasswordValid := auth.VerifyPassword([]byte(loginRequest.Password), []byte(user.Password))
	if isPasswordValid == false {
		app_http.SetJsonResponse(
			rw,
			dto.ErrorResponse{
				Error: httperrors.AUTH_INVALID_DETAILS,
			},
		)

		return
	}

	accessTokenExpiryTimestamp := time.Now().Add(time.Hour).Unix()
	refreshTokenExpiryTimestamp := time.Now().Add(time.Hour * 24).Unix()

	accessToken, accessTokenError := auth.GenerateJWT(
		strconv.FormatUint(uint64(user.ID), utils.DecimalDisplay),
		accessTokenExpiryTimestamp,
		[]byte(env.JwtSecret.GetValue()),
	)
	if accessTokenError != nil {
		a.Logger.Error(logserrors.AUTH_CONTROLLER_ERROR, "Access token generation error", accessTokenError)
		app_http.SetJsonResponse(
			rw,
			dto.ErrorResponse{
				Error: httperrors.API_GENERIC,
			},
		)

		return
	}

	refreshToken, refreshTokenError := auth.GenerateJWT(
		strconv.FormatUint(uint64(user.ID), utils.DecimalDisplay),
		refreshTokenExpiryTimestamp,
		[]byte(env.JwtSecret.GetValue()),
	)
	if refreshTokenError != nil {
		a.Logger.Error(logserrors.AUTH_CONTROLLER_ERROR, "Refresh token generation error", refreshTokenError)
		app_http.SetJsonResponse(
			rw,
			dto.ErrorResponse{
				Error: httperrors.API_GENERIC,
			},
		)

		return
	}

	a.DB.Model(&user).Updates(models.User{AccessToken: accessToken, RefreshToken: refreshToken})

	app_http.SetJsonResponse(
		rw,
		AuthResponse{
			UserId:       user.ID,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	)
}

func (a *AuthController) Logout(rw http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(middleware.JwtTokenContextKey).(*jwt.Token)
	subjectId, _ := token.Claims.GetSubject()

	a.DB.Model(&models.User{}).Where("id = ?", subjectId).Updates(map[string]any{
		"access_token":  nil,
		"refresh_token": nil,
	})

	parsedId, _ := strconv.ParseUint(subjectId, utils.DecimalDisplay, 0)

	app_http.SetJsonResponse(
		rw,
		AuthResponse{
			UserId: uint(parsedId),
		},
	)
}
