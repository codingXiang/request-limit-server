package server

import (
	"github.com/codingXiang/request-limit-server/util/backend"
	limit2 "github.com/codingXiang/request-limit-server/util/limit"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

const (
	CountKey = "_count"
)

func CheckIp(limit *viper.Viper, client *backend.RedisClient) gin.HandlerFunc {
	limitRange := limit2.NewRange(limit)
	return func(c *gin.Context) {
		var (
			clientIP = c.ClientIP()
		)
		if v := client.Get(c, clientIP); v.Err() != nil{
			client.Set(c, clientIP, "access", limitRange.Get())
			client.Set(c, clientIP + CountKey, 0, -1)
		}
		c.Next()
	}
}

//LimitByCount
func LimitByCount(limit *viper.Viper, client *backend.RedisClient) gin.HandlerFunc {
	requestLimit := limit.GetInt("limit.request")
	return func(c *gin.Context) {
		var (
			clientIP = c.ClientIP()
			count    = 0
			err      error
		)

		if count, err = client.AutoIncrement(c, clientIP + CountKey); err != nil {
			client.Set(c, clientIP + CountKey, count, -1)
		}
		c.Set(CountKey, count)
		if requestLimit >= count {
			c.Next()
			return
		} else {
			c.String(http.StatusTooManyRequests, "Error")
			c.Abort()
			return
		}
	}
}
