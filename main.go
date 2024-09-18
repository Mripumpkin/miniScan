package main

import (
	"miniScan/models"
	config "miniScan/utils/conf"
	logger "miniScan/utils/log"
)

func main() {
	// 初始化配置
	cfg := config.LoadConfigProvider()
	rdb := models.GetRedis(cfg, 1)
	rdb1 := models.GetRedis(cfg, 2)
	db := models.GetDB(cfg)
	log := logger.NewLogger(cfg)
	logP := logger.LogPortScan
	logD := logger.LogDomain
	logW := logger.LogWebInfo
	log.Error(rdb)
	logP.Error(rdb1)
	logD.Warn(db)
	logW.Error(cfg.GetString("run_level"))
}
