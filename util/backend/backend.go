package backend

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"strconv"
)

type RedisClient struct {
	*redis.Client
}

//NewRedisClient 新增 RedisClient Instance
func NewRedisClient(config *viper.Viper) (*RedisClient, error) {
	//讀取參數
	var (
		client   = new(RedisClient)
		url      = config.GetString("redis.url")
		port     = config.GetInt("redis.port")
		password = config.GetString("redis.password")
		db       = config.GetInt("redis.db")
	)

	//設定連線資訊
	option := &redis.Options{
		Addr: fmt.Sprintf("%s:%d", url, port),
		DB:   db,
	}
	//設定密碼
	if password != "" {
		option.Password = password
	}
	//建立 client
	client.Client = redis.NewClient(option)
	_, err := client.Ping(context.Background()).Result()

	return client, err
}

//AutoIncrement 自動新增 1 筆紀錄
func (client *RedisClient) AutoIncrement(c context.Context, key string) (int, error) {
	if v := client.Get(c, key); v.Err() == nil {
		count, err := strconv.Atoi(v.Val())
		if err != nil {
			return 1, err
		}
		count += 1
		updateVal := client.Set(c, key, count, -1)
		return count, updateVal.Err()
	} else {
		updateVal := client.Set(c, key, 1, -1)
		return 1, updateVal.Err()
	}
}
