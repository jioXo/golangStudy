package utils

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key") // 建议用更安全的密钥

// 生成 JWT
func GenerateToken(UserID uint) (string, error) {
	claims := jwt.MapClaims{
		"userId": UserID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// 解析和验证 JWT
func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
}

// ParseTokenFromRequest 从授权标头中提取并解析 JWT 令牌
func ParseTokenFromRequest(c *gin.Context) (*jwt.Token, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, jwt.ErrSignatureInvalid
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, jwt.ErrSignatureInvalid
	}
	return jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
}
