package component

import (
	"os"
	"time"
	"wx/app/model"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type IDTokenCustomClaims struct {
	User *model.User `json:"user"`
	jwt.StandardClaims
}

// 产生token的函数
func GenerateToken(u *model.User) (string, error) {
	unixTime := time.Now().Unix()
	tokenExp := unixTime + 3600*12 // 12 hour from current time

	claims := IDTokenCustomClaims{
		User: u,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  unixTime,
			ExpiresAt: tokenExp,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 验证token的函数
func ParseToken(token string) (*IDTokenCustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &IDTokenCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*IDTokenCustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
