package main

import (
	"flag"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"go-wsbackend/pkg/common"
	"go-wsbackend/pkg/database"
	h "go-wsbackend/pkg/http"
	"go-wsbackend/pkg/util"
	"path"
)

var (
	dataFile string
	useMysql bool
)

func init() {
	flag.StringVar(&dataFile, "f", path.Join("..", "data", "data.json"), "")
	flag.BoolVar(&useMysql, "d", true, "")
	flag.Parse()
}

func run() {
	url := util.DefaultGetEnvString("DB_URL", "127.0.0.1")
	username := util.DefaultGetEnvString("DB_USERNAME", "root")
	password := util.DefaultGetEnvString("DB_PASSWORD", "")
	dbName := util.DefaultGetEnvString("DB_NAME", "wsbackend")
	conf := &common.Config{
		DataFile: dataFile,
		UseMysql: useMysql,
		Mysql: struct {
			Url          string
			Username     string
			Password     string
			DataBaseName string
		}{Url: url, Username: username, Password: password, DataBaseName: dbName},
		WechatApps: map[string]common.WechatConfig{
			// 垃圾分类S
			"wx315cb91190a527ec": {
				Config: wechat.Config{
					AppID:          "wx315cb91190a527ec",
					AppSecret:      "",
					Token:          "",
					EncodingAESKey: "",
					Cache:          cache.NewMemory(),
				},
				Template: common.WechatTemplate{
					Pass: "x",
					Deny: "x",
				},
			},
		},
	}
	if useMysql {
		database.Init(conf)
	}
	r := h.Init(conf)
	r.Run()
}

func main() {
	//run()
	run()
}
