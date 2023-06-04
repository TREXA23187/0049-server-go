package main

import (
	"0049-server-go/core"
	"0049-server-go/flag"
	"0049-server-go/global"
	"0049-server-go/routers"
)

func main() {
	// Read Configuration File
	core.InitConf()
	// Initialize Log
	global.Log = core.InitLogger()
	// Connect to database
	global.DB = core.InitGorm()
	// Connect Redis
	global.Redis = core.ConnectRedis()

	// Command parameter binding
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("The server is running on: %s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
