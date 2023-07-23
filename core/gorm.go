package core

import (
	"0049-server-go/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Warnln("MySQL not configured, cancel gorm connection")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		// Display all SQL log information in the development environment
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error) // Print only error SQL log information
	}
	global.MysqlLog = logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Log.Fatalf(fmt.Sprintf("[%s] MySQL connection failed", dsn))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // Maximum number of idle connections
	sqlDB.SetMaxOpenConns(100)              // Maximum number of connections
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // The maximum lifetime of the connection, which cannot exceed the wait_timeout of MySQL

	global.Log.Infof("The mysql is running on: %s", global.Config.Mysql.Host)
	return db
}
