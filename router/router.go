package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sse_brussels_node/server"
	"time"
)

func SetUpRouter(s *server.Server) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许的域名，可以使用 "*" 允许所有
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // 预检请求的缓存时间
	}))

	r.GET("api/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("api/amm/order/status", s.GetTxStatus)

	r.GET("api/amm/order/list", s.GetOrderList)

	r.GET("api/amm/estimate", s.EstimateSwapResult)

	return r
}
