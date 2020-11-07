package main

import (
	limit_request "github.com/codingXiang/request-limit-server/module/limit-request"
	"github.com/codingXiang/request-limit-server/util/backend"
	"github.com/codingXiang/request-limit-server/util/config"
	"github.com/spf13/viper"
	"log"
)

var cm *viper.Viper
var limitCm *viper.Viper
var redisClient *backend.RedisClient

func init() {
	cm = config.Get(config.Yaml, "config", "./config")
	limitCm = config.Get(config.Yaml, "limit", "./config")

	backendCm := config.Get(config.Yaml, "backend", "./config")
	if err := limitCm.ReadInConfig(); err != nil {
		log.Fatalln("load config has failed, ", err)
	}
	if err := cm.ReadInConfig(); err != nil {
		log.Fatalln("load config has failed, ", err)
	}
	if err := backendCm.ReadInConfig(); err == nil {
		if client, err := backend.NewRedisClient(backendCm); err == nil {
			redisClient = client
		} else {
			log.Fatalln("redis connect failed, ", err)
		}
	} else {
		log.Fatalln("load config has failed, ", err)
	}
}
func main() {
	limit_request.NewLimitRequestServer(cm, nil).Init(limitCm, redisClient).Run()
}
