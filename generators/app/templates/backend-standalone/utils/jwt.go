package utils

import (
	"<%= displayName %>/global"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	UserId   uint   `json:"userId"`
	RoleID   uint   `json:"roleId"`
	jwt.StandardClaims
}

func CreateToken(username string, id uint, RoleID uint) (error, string) {

	claims := Claims{
		username,
		id,
		RoleID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 一周
			Issuer:    "jwt-issuer",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(global.CONFIG.JWT.SigningKey))

	return err, token
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.CONFIG.JWT.SigningKey), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
