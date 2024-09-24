package task

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

type TaskPayload struct {
	UUID          string   `json:"uuid"`
	Msg           string   `json:"msg"`
	Format        string   `json:"format"`
	ScanType      string   `json:"scan_type"`
	Plugins       []string `json:"plugins"`
	ExecutionTime int      `json:"execution_time"`
	Delay         int      `json:"delay"`
	Implement     bool     `json:"implement"`
}

func NewExampleTask(log *logrus.Logger, task *TaskPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(task)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %v", err)
	}
	return asynq.NewTask(task.UUID, payload), nil
}

func EnqueueTaskHandler(c *gin.Context, client *asynq.Client, log *logrus.Logger) {
	var payload TaskPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	task, err := NewExampleTask(log, &payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	// 任务入队
	info, err := client.Enqueue(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enqueue task"})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"task_id": info.ID,
		"queue":   info.Queue,
		"status":  "enqueued",
	})
}
