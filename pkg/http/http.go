package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/miniprogram"
	"github.com/silenceper/wechat/template"
	"go-wsbackend/pkg/common"
	"go-wsbackend/pkg/service"
	"time"
)

var (
	cf *common.Config
	db *gorm.DB
	wechatSrv *wechat.Wechat
	tpl *template.Template
	wxa *miniprogram.MiniProgram
)


func Init(c *common.Config) *gin.Engine {
	cf = c
	db = cf.DB
	wechatSrv = wechat.NewWechat(&cf.Wechat)
	tpl = wechatSrv.GetTemplate()
	wxa = wechatSrv.GetMiniProgram()

	service.Init(cf)

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
		v1.POST("/ws/feedback", userFeedback)
		v1.GET("/wechat/validate", wechatValidate)
		v1.GET("/wechat/code2session", code2session)
		v1.GET("/wechat/send/audit", sendAuditTemplate)
	}

	// 管理后台接口
	api := r.Group("/api")
	{
		api.GET("/waste", fetchWaste)
		api.GET("/waste/reload", reloadWaste)
		api.POST("/waste", addWaste)
		api.POST("/waste/:id", updateWaste)
	}

	return r
}
