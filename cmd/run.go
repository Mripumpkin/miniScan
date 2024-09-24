package cmd

import (
	"miniScan/pkg/http"

	config "miniScan/utils/conf"

	"github.com/hibiken/asynq"
	"github.com/qiniu/qmgo"
	"github.com/sirupsen/logrus"
)

func Run(cfg config.Provider, db *qmgo.Database, log *logrus.Logger) {
	redisAddr := cfg.GetString("redis.host") + ":" + cfg.GetString("redis.port")
	redisPasswd := cfg.GetString("redis.password")
	dbIndex := cfg.GetInt("redis.async_index")
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr, Password: redisPasswd, DB: dbIndex})
	defer client.Close()

	// 启动 HTTP (Gin) 服务
	http.StartHTTPServer(client, log)
}
