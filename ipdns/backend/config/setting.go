package config

import (
	"log"

	"github.com/spf13/viper"
)

func GetConfig(fileName, fileType, filePath string) (conf *viper.Viper, err error) {
	config := viper.New()
	config.SetConfigName(fileName)
	config.SetConfigType(fileType)
	config.AddConfigPath(filePath)

	err = config.ReadInConfig()
	if err != nil {
		log.Println("读取配置文件异常: ", fileName, err)
		return
	}

	return config, nil

}
