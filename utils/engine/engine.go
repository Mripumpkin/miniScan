package util

import (
	"context"
	"fmt"

	"miniScan/models/engine"
	config "miniScan/utils/conf"

	"github.com/docker/docker/api/types/container"
	client "github.com/docker/docker/client"
	"github.com/qiniu/qmgo"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func FindString(target string, intarray []string) bool {
	for _, element := range intarray {
		if target == element {
			return true
		}
	}
	return false
}

// DockerOperate 操作容器信息
func DockerOperate(cfg config.Provider, mongodb *qmgo.Database, logger *logrus.Logger) {

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		logger.Errorf("Failed To Get Docker Information:%v", err)
		return
	}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		logger.Errorf("Failed To Get Docker Information:%v", err)
		return
	}
	mongo := cfg.GetString("mongo.name")
	redis := cfg.GetString("redis.name")
	engines := []string{mongo, redis}
	DockerCollection := mongodb.Collection(cfg.GetString("mongo.table.engine_docker"))
	for _, container := range containers {
		Newcontainer := new(engine.DockerContainer)
		Newcontainer.Name = container.Names[0][1:]
		Newcontainer.CreateAt = container.Created
		Newcontainer.ContainerID = container.ID
		Newcontainer.ImageID = container.ImageID
		if len(container.Ports) > 0 {
			Newcontainer.Port = int16(container.Ports[0].PublicPort)
		}

		filter := bson.M{"name": Newcontainer.Name, "port": Newcontainer.Port}
		replacement := qmgo.M{
			"name":         Newcontainer.Name,
			"container_id": Newcontainer.ContainerID,
			"port":         Newcontainer.Port,
			"create_at":    Newcontainer.CreateAt,
			"image_id":     Newcontainer.ImageID,
		}

		if FindString(Newcontainer.Name, engines) {
			_, err := DockerCollection.Upsert(context.TODO(), filter, replacement)
			if err != nil {
				logger.Errorf("Failed To Get Docker Information:%v", err)
				return
			}
		}
	}
	logger.Info("Successed To Get Docker Information")
}

// 重启容器
func RestartContainer(containerID string) error {

	cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	err := cli.ContainerRestart(context.Background(), containerID, container.StopOptions{})
	if err != nil {
		return err
	}
	fmt.Println("Restart Container")
	return nil
}

// 启动容器
func StartContainer(containerID string) error {

	ctx := context.Background()
	cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	err := cli.ContainerStart(ctx, containerID, container.StartOptions{})
	if err != nil {
		return err
	}
	fmt.Println("Start Container")
	return nil
}

// 停止容器
func StopContainer(containerID string) error {

	ctx := context.Background()
	cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	err := cli.ContainerStop(ctx, containerID, container.StopOptions{})
	if err != nil {
		return err
	}
	fmt.Println("Stop Container")
	return nil
}
