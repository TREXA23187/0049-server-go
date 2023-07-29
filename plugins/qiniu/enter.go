package qiniu

import (
	"0049-server-go/configs"
	"0049-server-go/global"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"time"
)

func getToken(q configs.QiNiu) string {
	accessKey := q.AccessKey
	secretKey := q.SecretKey
	bucket := q.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

func getCfg(q configs.QiNiu) storage.Config {
	cfg := storage.Config{}

	zone, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
	cfg.Zone = &zone
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false
	return cfg

}

func UploadFile(data []byte, fileName string, prefix string) (filePath string, err error) {
	if !global.Config.QiNiu.Enable {
		return "", errors.New("please enable oss")
	}

	q := global.Config.QiNiu
	if q.AccessKey == "" || q.SecretKey == "" {
		return "", errors.New("please config accessKey and secretKey")
	}
	if float64(len(data))/1024/1024 > q.Size {
		return "", errors.New("the file exceeds the maximum size")
	}
	upToken := getToken(q)
	cfg := getCfg(q)

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	dataLen := int64(len(data))

	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf("%s/%s_%s", prefix, now, fileName)

	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", q.CDN, ret.Key), nil
}
