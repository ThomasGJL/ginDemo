package gin

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"ginDemo/config/log"
	err "ginDemo/exception"
)

// InitGinConfig 初始化Gin
func InitGinConfig() *gin.Engine {
	logger.Info("Gin: Init gin start")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// 入口日志打印
	router.Use(log.LoggerAccess)
	// 统一异常处理
	router.Use(err.ErrHandle)
	// 跨域处理
	router.Use(cors.Default())
	// token校验
	//router.Use(token.TokenVerify)
	// 健康检测
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server up and running",
		})
	})
	logger.Info("Gin: Init gin done")
	return router
}

// RunGin 启动Gin
func RunGin(router *gin.Engine) {
	port := viper.GetString("server.port")
	logger.Info(fmt.Sprintf("Service started on port(s): %s", port))
	_ = router.Run(":" + port)
}
