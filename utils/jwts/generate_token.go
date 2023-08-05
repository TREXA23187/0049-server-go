package jwts

import (
	"0049-server-go/global"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

func GenerateToken(user JwtPayLoad) (string, error) {

	MySecret = []byte(global.Config.Jwt.Secret)

	claim := CustomClaims{
		JwtPayLoad: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))), // Default 2 hour expiration
			Issuer:    global.Config.Jwt.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString(MySecret)
}

func GenerateLinkToken(link LinkJwtPayLoad, expires time.Duration) (string, error) {

	MySecret = []byte(global.Config.Jwt.Secret)

	claim := LinkCustomClaims{
		LinkJwtPayLoad: link,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))), // Default 2 hour expiration
			Issuer:    global.Config.Jwt.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString(MySecret)
}
