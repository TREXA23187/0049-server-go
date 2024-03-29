package jwts

import (
	"0049-server-go/global"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
)

func ParseToken(tokenStr string) (*CustomClaims, error) {
	MySecret = []byte(global.Config.Jwt.Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		global.Log.Error(fmt.Sprintf("token parse err: %s", err.Error()))
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func ParseLinkToken(tokenStr string) (*LinkCustomClaims, error) {
	MySecret = []byte(global.Config.Jwt.Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &LinkCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		global.Log.Error(fmt.Sprintf("link token parse err: %s", err.Error()))
		return nil, err
	}
	if claims, ok := token.Claims.(*LinkCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid link token")
}
