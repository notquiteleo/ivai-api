package jwt

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your_secret_key") // 应该使用环境变量或安全的方式存储

// CustomClaims 包含JWT的声明
type CustomClaims struct {
	UserID string `json:"userId"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT
func GenerateToken(userID string) (string, error) {
	// 设置JWT的声明
	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 设置过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 设置签发时间
		},
	}

	// 创建一个新的JWT，并使用HS256算法进行签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用密钥签名并获取完整的编码后的字符串表示形式
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken 校验JWT
func ValidateToken(tokenString string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	// 解析JWT字符串，并存储结果在claims中
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func JWTMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
				return
			}

			_, err := ValidateToken(token)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
