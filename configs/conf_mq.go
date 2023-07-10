package configs

import (
	"fmt"
)

type MQ struct {
	IP       string `json:"ip" yaml:"ip"`
	Port     int    `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

func (mq MQ) Addr() string {
	return fmt.Sprintf("%s:%d", mq.IP, mq.Port)
}

func (mq MQ) URL() string {
	return fmt.Sprintf("amqp://%s:%s@%s/", mq.Username, mq.Password, mq.Addr())
}
