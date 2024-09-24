package main

import (
	cmd "miniScan/cmd"
	"miniScan/models"
	config "miniScan/utils/conf"
	engine "miniScan/utils/engine"
	logger "miniScan/utils/log"
	"runtime/debug"
)

func main() {
	// 初始化配置
	cfg := config.LoadConfigProvider()
	db := models.GetDB(cfg)
	log := logger.NewLogger(cfg)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				s := string(debug.Stack())
				log.Errorf("err=%v, stack=%s\n", err, s)
			}
		}()
	}()

	go engine.DockerOperate(cfg, db, log)
	// 启动 HTTP (Gin) 服务
	go cmd.Run(cfg, db, log)
	select {}
}
