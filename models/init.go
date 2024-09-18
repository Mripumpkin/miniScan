package models

import (
	"context"
	"fmt"
	"net/url"
	"sync"
	"time"

	config "miniScan/utils/conf"

	"github.com/go-redis/redis/v8"
	"github.com/qiniu/qmgo"
)

var onceDB sync.Once
var redisClients sync.Map

func GetDB(cfg config.Provider) *qmgo.Database {
	var db *qmgo.Database
	onceDB.Do(func() {
		db = initDB(cfg)
	})
	return db
}

func GetRedis(cfg config.Provider, dbIndex int) *redis.Client {
	client, ok := redisClients.Load(dbIndex)
	if ok {
		return client.(*redis.Client)
	}

	newClient := GetRedisClient(cfg, dbIndex)
	redisClients.Store(dbIndex, newClient)
	return newClient
}

func initDB(cfg config.Provider) *qmgo.Database {
	ctx := context.Background()
	username := cfg.GetString("mongo.username")
	password := cfg.GetString("mongo.password")
	host := cfg.GetString("mongo.host")
	port := cfg.GetString("mongo.port")
	dbname := cfg.GetString("mongo.dbname")
	URI := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin&directConnection=true", username, url.QueryEscape(password), host, port)
	timeout := cfg.GetInt64("mongo.conn_timeout")
	client, err := qmgo.NewClient(ctx, &qmgo.Config{
		Uri:              URI,
		ConnectTimeoutMS: &timeout,
	})
	if err != nil {
		panic(fmt.Sprintf("db conn err: %v", err))
	}

	err = client.Ping(5)
	if err != nil {
		panic(fmt.Sprintf("db ping err: %v", err))
	}

	return client.Database(dbname)
}

var ctx = context.Background()

func GetRedisClient(cfg config.Provider, dbIndex int) *redis.Client {
	host := fmt.Sprintf("%v:%v", cfg.GetString("redis.host"), cfg.GetString("redis.port"))
	password := cfg.GetString("redis.password")
	connTimeout := time.Duration(cfg.GetInt64("redis.conn_timeout"))

	rdb := redis.NewClient(&redis.Options{
		Addr:         host,
		Password:     password,
		DB:           dbIndex,
		PoolSize:     15,
		MinIdleConns: 1,
		DialTimeout:  time.Millisecond * connTimeout,
		ReadTimeout:  time.Millisecond * connTimeout,
		WriteTimeout: time.Millisecond * connTimeout,
	})

	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic(fmt.Sprintf("rdb ping err: %v", err))
	}

	return rdb
}

// // 初始化ES连接
// func ElasticInit(urls []string, name string, pwd string, cfg *viper.Viper) *elastic.Client {
// 	var client *elastic.Client
// 	var err error
// 	if name != "" && pwd != "" {
// 		client, err = elastic.NewSimpleClient(elastic.SetURL(urls...), elastic.SetBasicAuth(name, pwd), elastic.SetSniff(true))
// 	} else {
// 		client, err = elastic.NewSimpleClient(elastic.SetURL(urls...), elastic.SetSniff(true))
// 	}
// 	if err != nil {
// 		fmt.Printf("Elastic client init ERROR: %v", err)
// 	}
// 	return client
// }
