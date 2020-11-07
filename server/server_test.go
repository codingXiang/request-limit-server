package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type ServerSuite struct {
	suite.Suite
	config *viper.Viper
	engine *gin.Engine
	server *Server
}

func (s *ServerSuite) SetupSuite() {
	config := viper.New()
	config.Set("application.mode", "release")
	config.Set("application.timeout.read", 1000)
	config.Set("application.timeout.write", 1000)
	config.Set("application.port", 8888)
	s.config = config
}

func (s *ServerSuite) BeforeTest(suiteName, testName string) {
	s.server = New().Init(s.config, nil)
}

//TestStart 為測試程式進入點
func TestStartServerSuite(t *testing.T) {
	suite.Run(t, new(ServerSuite))
}

func (s *ServerSuite) TestNew() {
	server := New()
	assert.Equal(s.T(), server, new(Server))
}

func (s *ServerSuite) TestInit() {
	{
		s.server = New().Init(s.config, nil)
		ser := s.server.server
		assert.Equal(s.T(), ser.Addr, ":8888", "server port is not correct")
		assert.Equal(s.T(), ser.ReadTimeout, time.Duration(1000) * time.Second)
		assert.Equal(s.T(), ser.WriteTimeout, time.Duration(1000) * time.Second)
	}
	{
		eng := gin.Default()
		s.server = New().Init(s.config, eng)
		ser := s.server.server
		assert.Equal(s.T(), ser.Addr, ":8888", "server port is not correct")
		assert.Equal(s.T(), ser.Handler, eng)
		assert.Equal(s.T(), ser.ReadTimeout, time.Duration(1000) * time.Second)
		assert.Equal(s.T(), ser.WriteTimeout, time.Duration(1000) * time.Second)
	}
	assert.Equal(s.T(), s.server.config, s.config)
}

func (s *ServerSuite) TestGetEngine() {
	eng := gin.Default()
	s.server = New().Init(s.config, eng)
	assert.Equal(s.T(), s.server.GetEngine(), eng)
}

func (s *ServerSuite) TestGetServer() {
	{
		s.server = New().Init(s.config, nil)
		assert.NotNil(s.T(), s.server.GetServer())
	}
}
