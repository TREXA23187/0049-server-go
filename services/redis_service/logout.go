package redis_service

import (
	"0049-server-go/global"
	"0049-server-go/utils"
	"fmt"
	"time"
)

type RedisService struct {
}

const logoutPrefix = "logout_"

// Logout For logout operations
func Logout(token string, diff time.Duration) error {
	err := global.Redis.Set(fmt.Sprintf("%s%s", logoutPrefix, token), "", diff).Err()
	return err
}

func CheckLogout(token string) bool {
	keys := global.Redis.Keys(logoutPrefix + "*").Val()
	if utils.InList(logoutPrefix+token, keys) {
		return true
	}

	return false
}
