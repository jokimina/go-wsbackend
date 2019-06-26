package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-wsbackend/pkg/common"
	"time"
)

var cf *common.Config


func Init(c *common.Config) *gin.Engine{
	cf = c
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
	r.GET("/v1/ws", getAllWaste)
	return r
}

