package log_stash

import (
	"0049-server-go/global"
	"0049-server-go/utils/jwts"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Log struct {
	Ip     string `json:"ip"`
	Addr   string `json:"addr"`
	UserId uint   `json:"user_id"`
}

func New(ip string, token string) *Log {
	// 解析token
	claims, err := jwts.ParseToken(token)
	var userID uint
	if err == nil {
		userID = claims.UserID
	}

	// 拿到用户id
	return &Log{
		Ip:     ip,
		Addr:   "内网",
		UserId: userID,
	}
}

func NewLogByGin(c *gin.Context) *Log {
	ip := c.ClientIP()
	token := c.Request.Header.Get("token")
	return New(ip, token)
}

func (l Log) Debug(content string) {
	l.send(DebugLevel, content)
}
func (l Log) Info(content string) {
	l.send(InfoLevel, content)
}
func (l Log) Warn(content string) {
	l.send(WarningLevel, content)
}
func (l Log) Error(content string) {
	l.send(ErrorLevel, content)
}

func (l Log) send(level Level, content string) {
	err := global.DB.Create(&LogStashModel{
		IP:      l.Ip,
		Addr:    l.Addr,
		Level:   level,
		Content: content,
		UserID:  l.UserId,
	}).Error
	if err != nil {
		logrus.Error(err)
	}
}
