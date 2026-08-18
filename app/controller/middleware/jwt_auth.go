package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"api-backend/sample/app/application/env"
	"api-backend/sample/app/controller/dto"
	"api-backend/sample/app/service/auth"
	app_http "api-backend/sample/app/service/http"
)

type JwtTokenContextKeyType string

const JwtTokenContextKey JwtTokenContextKeyType = "JWT_TOKEN_KEY"

func JwtAuthCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		if authorization == "" {
			app_http.SetJsonResponse(
				w,
				dto.ErrorResponse{
					Error: "Incorrect authorization header",
				},
			)

			return
		}

		authorizationSlice := strings.Split(authorization, " ")
		if len(authorizationSlice) != 2 || authorizationSlice[0] != "Bearer" {
			app_http.SetJsonResponse(
				w,
				dto.ErrorResponse{
					Error: "Incorrect authorization header",
				},
			)

			return
		}

		token, err := auth.VerifyJWT(authorizationSlice[1], []byte(env.JwtSecret.GetValue()))
		if err != nil {
			app_http.SetJsonResponse(
				w,
				dto.ErrorResponse{
					Error: "Invalid JWT",
				},
			)

			return
		}

		expiresAt, err := token.Claims.GetExpirationTime()
		if err != nil {
			app_http.SetJsonResponse(
				w,
				dto.ErrorResponse{
					Error: "Invalid JWT",
				},
			)
		}

		if expiresAt.Unix() < time.Now().Unix() {
			app_http.SetJsonResponse(
				w,
				dto.ErrorResponse{
					Error: "Invalid JWT",
				},
			)

			return
		}

		newCtx := context.WithValue(r.Context(), JwtTokenContextKey, token)
		r = r.WithContext(newCtx)

		next.ServeHTTP(w, r)
	})
}
