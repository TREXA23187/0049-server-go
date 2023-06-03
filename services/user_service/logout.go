package user_service

import (
	"0049-server-go/services/redis_service"
	"0049-server-go/utils/jwts"
	"time"
)

func (UserService) Logout(claims *jwts.CustomClaims, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)

	return redis_service.Logout(token, diff)
}
