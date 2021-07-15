package component

import (
	"os"
	"time"
	"wx/app/dto"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cast"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
var jwtExpSec = os.Getenv("JWT_EXP_SEC")

type IDTokenCustomClaims struct {
	User *dto.UserJWT `json:"user"`
	jwt.StandardClaims
}

// 产生token的函数
func GenerateToken(u *dto.UserJWT) (string, error) {
	unixTime := time.Now().Unix()
	tokenExp := unixTime + cast.ToInt64(jwtExpSec)
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
