package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-wsbackend/pkg/common"
	"go-wsbackend/pkg/service"
	"time"
)

var (
	cf *common.Config
	db *gorm.DB
)


func Init(c *common.Config) *gin.Engine {
	cf = c
	service.Init(cf)
	db = cf.DB
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())

	// 提供小程序端访问接口
	v1 := r.Group("/v1")
	{
		v1.GET("/ws", getAllWaste)
	}

	api := r.Group("/api")
	{
		api.GET("/waste", fetchWaste)
		api.POST("/waste", addWaste)
		api.POST("/waste/:id", updateWaste)
	}

	return r
}
