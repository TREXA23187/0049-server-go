package redis_service

import (
	"0049-server-go/global"
	"0049-server-go/utils"
	"fmt"
	"time"
)

const buildImagePrefix = "build_image_"

func StartBuildingImage(repository string, diff time.Duration) error {
	err := global.Redis.Set(fmt.Sprintf("%s%s", buildImagePrefix, repository), "", diff).Err()
	return err
}

func FinishBuildingImage(repository string) error {
	err := global.Redis.Del(fmt.Sprintf("%s%s", buildImagePrefix, repository)).Err()
	return err
}

func CheckIsImageInBuilding(repository string) bool {
	keys := global.Redis.Keys(buildImagePrefix + "*").Val()
	if utils.InList(buildImagePrefix+repository, keys) {
		return true
	}

	return false
}
