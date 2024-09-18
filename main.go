package main

import (
	"fmt"
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
				fmt.Printf("err=%v, stack=%s\n", err, s)
			}
		}()
	}()
	// rdb := models.GetRedis(cfg, 1)
	// rdb1 := models.GetRedis(cfg, 2)

	// logP := logger.LogPortScan
	// logD := logger.LogDomain
	// logW := logger.LogWebInfo
	// log.Error(rdb)
	// logP.Error(rdb1)
	// logD.Warn(db)
	// logW.Error(cfg.GetString("run_level"))
	go engine.DockerOperate(cfg, db, log)
}
