package common

import "github.com/jinzhu/gorm"

type Config struct {
	DataFile string
	DB	*gorm.DB
	UseMysql bool
	Mysql struct{
		Url string
		Username string
		Password string
		DataBaseName string
	}
}


