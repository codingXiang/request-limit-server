package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type Server struct {
	config *viper.Viper
	engine *gin.Engine
	server *http.Server
}

func New() *Server {
	return new(Server)
}

func (s *Server) Init(config *viper.Viper, engine *gin.Engine) *Server {
	//設定 gin 啟動模式
	gin.SetMode(config.GetString("application.mode"))
	//設定 server config
	s.config = config
	//設定 server engine
	if engine == nil {
		s.engine = gin.Default()
	} else {
		s.engine = engine
	}
	var (
		port         = config.GetInt("application.port")          //伺服器的 port
		writeTimeout = config.GetInt("application.timeout.write") //伺服器的寫入超時時間
		readTimeout  = config.GetInt("application.timeout.read")  //伺服器讀取超時時間
	)
	// 設定 http server
	s.server = &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        s.engine,
		ReadTimeout:    time.Duration(readTimeout) * time.Second,
		WriteTimeout:   time.Duration(writeTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s
}

func (s *Server) GetEngine() *gin.Engine {
	return s.engine
}

func (s *Server) GetServer() *http.Server {
	return s.server
}