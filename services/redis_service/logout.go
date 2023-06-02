package redis_service

import (
	"0049-server-go/global"
	"0049-server-go/utils"
	"fmt"
	"time"
)

type RedisService struct {
}

const prefix = "logout_"

// Logout For logout operations
func Logout(token string, diff time.Duration) error {
	err := global.Redis.Set(fmt.Sprintf("%s%s", prefix, token), "", diff).Err()
	return err
}

func CheckLogout(token string) bool {
	keys := global.Redis.Keys(prefix + "*").Val()
	if utils.InList(prefix+token, keys) {
		return true
	}

	return false
}
