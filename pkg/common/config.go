package common

import "github.com/jinzhu/gorm"

type Config struct {
	DataFile string
	DB	*gorm.DB
}


