package common

import (
	"github.com/jinzhu/gorm"
	"github.com/silenceper/wechat"
)

type Config struct {
	DataFile string
	DB	*gorm.DB
	UseMysql bool
	WechatApps map[string]wechat.Config
	Mysql struct{
		Url string
		Username string
		Password string
		DataBaseName string
	}
}


