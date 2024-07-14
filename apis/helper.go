package apis

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Response struct {
	Success bool
	Message string
	Data    any
}

func RespondWith(w http.ResponseWriter, r *http.Request, route string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		RespondWithError(w, r, route, err, http.StatusInternalServerError)
		return
	}
}

func RespondWithError(w http.ResponseWriter, r *http.Request, route string, err error, responseCode int) {
	type response struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
	}

	if errors.Is(err, context.DeadlineExceeded) {
		responseCode = http.StatusGatewayTimeout
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)

	resp := response{
		Success: false,
		Error:   err.Error(),
	}

	_ = json.NewEncoder(w).Encode(resp) //nolint:errchkjson
}

var jwtKey = []byte("your_secret_key")

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userID int64) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 设置令牌有效期为24小时
	claims := &Claims{
			UserID: userID,
			StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
			},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}