package http

import (
	"miniScan/pkg/task"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

func StartHTTPServer(client *asynq.Client, log *logrus.Logger) {
	r := gin.Default()

	r.POST("/enqueue", func(c *gin.Context) {
		task.EnqueueTaskHandler(c, client, log)
	})

	log.Println("Starting Gin server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Gin server failed: %v", err)
	}
}
