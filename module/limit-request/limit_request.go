package limit_request

import (
	"github.com/codingXiang/request-limit-server/server"
	"github.com/codingXiang/request-limit-server/util/backend"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)


type LimitRequestServer struct {
	*server.Server
}

func NewLimitRequestServer(config *viper.Viper, engine *gin.Engine) *LimitRequestServer {
	s := new(LimitRequestServer)
	s.Server = server.New().Init(config, engine)
	return s
}

//Init 初始化 Limit Request Server
func (s *LimitRequestServer) Init(config *viper.Viper, client *backend.RedisClient) *LimitRequestServer {
	s.Server.GetEngine().
		//先判斷 ip 存取
		Use(server.CheckIp(config, client)).
		//接著判斷存取次數是否超出設定值
		Use(server.LimitByCount(config, client)).
		GET("/", s.Index)
	return s
}

//Index 首頁，顯示目前存取次數
func (s *LimitRequestServer) Index(c *gin.Context) {
	c.String(http.StatusOK, "%d", c.GetInt(server.CountKey))
}

//Run 運行 Server
func (s *LimitRequestServer) Run() {
	s.Server.GetServer().ListenAndServe()
}
