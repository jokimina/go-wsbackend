package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/silenceper/wechat"
	"go-wsbackend/pkg/common"
	"go-wsbackend/pkg/service"
	"go-wsbackend/pkg/util"
	"io"
	"os"
	"path"
	"time"
)

var (
	cf *common.Config
	db *gorm.DB
	apps map[string]*wechat.Wechat
	wechatSrv *wechat.Wechat
	logDir = path.Join("..", "logs")
	logFile = path.Join(logDir, "wsbackend.log")
	//tpl *template.Template
	//wxa *miniprogram.MiniProgram
)


func Init(c *common.Config) *gin.Engine {
	apps = make(map[string]*wechat.Wechat)
	cf = c
	db = cf.DB
	for appId, wc := range cf.WechatApps {
		apps[appId] = wechat.NewWechat(&wc.Config)
	}

	service.Init(cf)

	r := gin.New()
	f := logFileInit()
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

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
	r.Use(decryptMiddleware())

	// 提供小程序端访问接口
	v1 := r.Group("/v1")
	{
		v1.GET("/ws", getAllWaste)
		v1.GET("/ws/count", getWasteCount)
		v1.GET("/ws/search", searchWaste)
		v1.GET("/ws/checksum", getWasteChecksum)
		v1.POST("/ws/feedback", userFeedback)
		v1.POST("/wechat/push", wechatPush)
		v1.GET("/wechat/code2session", code2session)
		v1.GET("/wechat/send/audit", sendAuditTemplate)
		v1.POST("/wechat/userinfo", saveUserInfo)
	}

	// 管理后台接口
	api := r.Group("/api")
	{
		api.GET("/waste", fetchWaste)
		api.GET("/waste/reload", reloadWaste)
		api.POST("/waste", addWaste)
		api.POST("/waste/:id", updateWaste)
		api.POST("/audit/waste", auditWaste)
		api.POST("/audit/batch/waste", auditBatchWaste)
	}

	return r
}


func logFileInit()(f *os.File){
	exist, err := util.PathExists(logDir)
	if err != nil {
		panic(err)
	}
	if !exist {
		if err = os.Mkdir(logDir, os.ModePerm); err != nil {
			panic(err)
		}
	}
	f, err = os.Create(logFile)
	if err != nil {
		panic(err)
	}
	return
}