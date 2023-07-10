package global

import (
	"0049-server-go/configs"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Config   *configs.Config
	DB       *gorm.DB
	Log      *logrus.Logger
	MysqlLog logger.Interface
	Redis    *redis.Client
	MQ       *amqp.Connection
)
