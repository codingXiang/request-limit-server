package backend

import (
	"context"
	"github.com/alicebob/miniredis/v2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)


//Suite struct
type Suite struct {
	suite.Suite
	mockRedis *miniredis.Miniredis
	config    *viper.Viper
	client    *RedisClient
}

//初始化 Suite
func (s *Suite) SetupSuite() {
	const (
		DB       = 0
		Password = "a12345"
	)
	mock, _ := miniredis.Run()
	mock.RequireAuth(Password)
	s.mockRedis = mock
	s.config = viper.New()
	s.config.Set("redis.url", s.mockRedis.Host())
	s.config.Set("redis.port", s.mockRedis.Port())
	s.config.Set("redis.db", DB)
	s.config.Set("redis.password", Password)

}

func (s *Suite) BeforeTest(suiteName, testName string) {
	s.client, _ = NewRedisClient(s.config)
}

//TestStart 為測試程式進入點
func TestStart(t *testing.T) {
	suite.Run(t, new(Suite))
}

//TestNewRedisClient 用於測試 NewRedisClient
func (s *Suite) TestNewRedisClient() {
	var (
		err error
		c   = context.TODO()
	)
	s.client, err = NewRedisClient(s.config)
	assert.Nil(s.T(), err)
	s.client.Set(c, "test", "test", -1)
	res, err := s.mockRedis.Get("test")
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "test", res)
}

//TestAutoIncrement 用於測試 AutoIncrement
func (s *Suite) TestAutoIncrement() {
	count, err := s.client.AutoIncrement(context.TODO(), "test")
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 1, count)
	count, err = s.client.AutoIncrement(context.TODO(), "test")
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 2, count)

}
