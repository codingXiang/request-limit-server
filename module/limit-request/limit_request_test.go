package limit_request

import (
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"github.com/codingXiang/request-limit-server/util/backend"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Suite struct {
	suite.Suite
	config  *viper.Viper
	client  *backend.RedisClient
	engine  *gin.Engine
	resp    *httptest.ResponseRecorder
	context *gin.Context
	server  *LimitRequestServer
}

func (s *Suite) SetupSuite() {
	config := viper.New()
	config.Set("limit.range.unit", "minute")
	config.Set("limit.range.per", "1")
	config.Set("limit.request", "60")
	s.config = config

	const (
		DB       = 0
		Password = "a12345"
	)
	mock, _ := miniredis.Run()
	mock.RequireAuth(Password)
	s.config.Set("redis.url", mock.Host())
	s.config.Set("redis.port", mock.Port())
	s.config.Set("redis.db", DB)
	s.config.Set("redis.password", Password)

	s.client, _ = backend.NewRedisClient(s.config)
}

func (s *Suite) BeforeTest(suiteName, testName string) {
	s.resp = httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	s.context, s.engine = gin.CreateTestContext(s.resp)
	s.server = NewLimitRequestServer(s.config, s.engine)
}

func (s *Suite) createRequest() {
	s.resp = httptest.NewRecorder()
	s.context, _ = gin.CreateTestContext(s.resp)
	s.context.Request = nil
	s.context.Request, _ = http.NewRequest("GET", "/", nil)
	s.context.Request.Header.Set("X-Forwarded-For", "127.0.0.1")
	s.engine.ServeHTTP(s.resp, s.context.Request)
}

//TestStart 為測試程式進入點
func TestStartSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestNewLimitRequestServer() {
	assert.NotNil(s.T(), s.server)
	assert.Equal(s.T(), s.engine, s.server.GetEngine())
}

func (s *Suite) TestIndex() {
	s.server.GetEngine().GET("/", s.server.Index)
	s.createRequest()
	assert.Equal(s.T(), http.StatusOK, s.resp.Code)
}

func (s *Suite) TestInit() {
	s.server.Init(s.config, s.client)
	// access
	for i := 1; i <= 60; i++ {
		s.createRequest()
		assert.Equal(s.T(), 200, s.resp.Code)
		assert.Equal(s.T(), s.context.ClientIP(), "127.0.0.1")
		actual := s.resp.Body.String()
		assert.Equal(s.T(), actual, fmt.Sprintf("%d", i))
	}
	// error
	for i := 1; i < 5; i++ {
		s.createRequest()
		assert.Equal(s.T(), http.StatusTooManyRequests, s.resp.Code)
		actual := s.resp.Body.String()
		assert.Equal(s.T(), actual, "Error")
	}
}
