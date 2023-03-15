package cmd

import (
	"file_flow/config"
	"file_flow/global"
	"file_flow/load"
	"file_flow/router"
	"fmt"
)

func Start() {
	config.InitConfig()
	global.Client = load.InitDB()
	global.Redis = load.InitRedis()
	global.Minio = load.InitMinio()
	router.InitRoute()
}

func Clean() {
	fmt.Println("Clean....")
	err := global.Client.Close()
	if err != nil {
		return
	}
	err = global.Redis.Close()
}
