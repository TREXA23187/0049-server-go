package main

import (
	"0049-server-go/core"
	"0049-server-go/flag"
	"0049-server-go/global"
	"0049-server-go/routers"
	"embed"
)

//go:embed settings.yaml
var embedConfigFiles embed.FS

// CGO_ENABLED=0 GOOS=linux go build -v -o server
func main() {
	// Read Configuration File
	core.InitConf(embedConfigFiles)
	// Initialize Log
	global.Log = core.InitLogger()
	// Connect to database
	global.DB = core.InitGorm()

	// Command parameter binding
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	// Connect Redis
	global.Redis = core.ConnectRedis()
	// Connect gRPC
	global.GRPC = core.ConnectGRPC()
	// Connect RabbitMQ
	global.MQ = core.ConnectRabbitMQ()

	go core.InitRabbitMQConsuming()

	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("The server is running on: %s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
