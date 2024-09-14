package main

import (
	config "miniScan/utils/conf"
	"miniScan/utils/log"
)

func main() {
	cfg := config.LoadConfigProvider()
	logger := log.NewLogger(cfg)
	logger.Error("---------------------")
}
