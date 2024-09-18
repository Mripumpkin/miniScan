package engine

type DockerContainer struct {
	Name        string `json:"name" bson:"name"`
	ImageID     string `json:"image_id" bson:"image_id"`
	ContainerID string `json:"Container_ID" bson:"container_id"`
	Port        int16  `json:"prot" bson:"port"`
	CreateAt    int64  `json:"create_at" bson:"create_at"`
}
