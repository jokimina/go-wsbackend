package common

import (
	"github.com/jinzhu/gorm"
	"github.com/silenceper/wechat"
)

type WechatTemplate struct {
	Pass string // 审核通过
	Deny string	 // 审核拒绝
}

type WechatConfig struct {
	Config wechat.Config
	Template WechatTemplate
}

type Config struct {
	DataFile string
	DB	*gorm.DB
	UseMysql bool
	WechatApps map[string]WechatConfig
	Mysql struct{
		Url string
		Username string
		Password string
		DataBaseName string
	}
}


