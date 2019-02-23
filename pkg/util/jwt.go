package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"jokeclub/pkg/setting"
	"time"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Claims struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(userId int, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * 365 * time.Hour)

	claims := Claims{
		userId,
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "jokeclub",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func GetToken(c *gin.Context) string {
	token := c.Query("token")
	if token == "" {
		token = c.PostForm("token")
	}
	return token
}
