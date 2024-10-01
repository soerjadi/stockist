package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/soerjadi/stockist/internal/config"
	"github.com/soerjadi/stockist/internal/delivery/rest"
	"github.com/soerjadi/stockist/internal/model/constant"
	"github.com/soerjadi/stockist/internal/pkg/log"
	"github.com/soerjadi/stockist/internal/usecase/user"
)

func OnlyLoggedInUser(userManagement user.Usecase, cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get token from header
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				log.Error("[middleware] token empty")
				unauthorizedResp(w, r)
				return
			}
			if !strings.Contains(tokenString, "Bearer") {
				log.Error("[middleware] bearer empty")
				unauthorizedResp(w, r)
				return
			}
			tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

			// Parse token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("signing method is invalid")
				} else if method != jwt.SigningMethodHS256 {
					return nil, errors.New("signing method is invalid")
				}

				return []byte(cfg.Secret.Token), nil
			})
			if err != nil {
				unauthorizedResp(w, r)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok && !token.Valid {
				unauthorizedResp(w, r)
				return
			}

			userIDStr, ok := claims["sub"].(string)
			if !ok {
				unauthorizedResp(w, r)
				return
			}

			userID, err := strconv.ParseInt(userIDStr, 10, 64)
			if err != nil {
				unauthorizedResp(w, r)
				return
			}

			user, err := userManagement.GetByID(context.Background(), userID)
			if err != nil {
				unauthorizedResp(w, r)
				return
			}

			reqMap := make(map[string]interface{})
			reqMap["user"] = user
			reqMap["userID"] = user.ID
			reqMap["req"] = r

			r = appendToContext(r.Context(), reqMap)

			next.ServeHTTP(w, r)
		})
	}
}

func unauthorizedResp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := rest.Response{
		Message: "Unauthenticated",
	}

	w.WriteHeader(http.StatusUnauthorized)
	x, _ := json.Marshal(resp)
	w.Write(x)
}

func appendToContext(ctx context.Context, reqMap map[string]interface{}) *http.Request {
	user := reqMap["user"]
	userID := reqMap["userID"]
	r := reqMap["req"].(*http.Request)
	ctx = context.WithValue(ctx, constant.USER_KEY_RESPONDENT, user)
	ctx = context.WithValue(ctx, constant.USER_ID_KEY_RESPONDENT, userID)

	r = r.WithContext(ctx)
	return r
}
