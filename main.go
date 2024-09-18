package main

import (
	config "miniScan/utils/conf"
	log "miniScan/utils/log"
)

func main() {
	// 初始化配置
	cfg := config.LoadConfigProvider() // 加载配置文件

	// 记录端口扫描日志
	logP := log.LogPortScan
	// 记录子域名扫描的警告日志
	logD := log.LogDomain
	// 记录 web 探测相关的信息
	logW := log.LogWebInfo
	logD.Info(cfg.GetBool(""))
	logP.Warn(cfg.GetBool(""))
	logW.Error(cfg.GetBool(""))
}
