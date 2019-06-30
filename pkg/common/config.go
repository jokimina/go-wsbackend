package common

import (
	"github.com/jinzhu/gorm"
	"github.com/silenceper/wechat"
)

type Config struct {
	DataFile string
	DB	*gorm.DB
	UseMysql bool
	Wechat wechat.Config
	Mysql struct{
		Url string
		Username string
		Password string
		DataBaseName string
	}
}


