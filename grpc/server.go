package main

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
	"google.golang.org/grpc"
)

var redisAddr = "127.0.0.1:6379"
var grpcAddr = "localhost:50051"

// ProcessTask is the worker that dequeues tasks
func ProcessTask(ctx context.Context, task *asynq.Task) error {
	// Connect to gRPC service
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewTaskExecutorClient(conn)
	req := &TaskP{
		TaskId:  "task_id_example",
		Payload: string(task.Payload()),
	}

	_, err = client.ExecuteTask(context.Background(), req)
	if err != nil {
		log.Fatalf("could not execute task: %v", err)
	}

	return nil
}

func main() {
	// Initialize Redis connection
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			Concurrency: 10, 
		},
	)

	// Register the task handler
	mux := asynq.NewServeMux()
	mux.HandleFunc("task_type", ProcessTask)

	// Start processing tasks
	if err := server.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
