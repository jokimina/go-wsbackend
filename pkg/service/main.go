package service

import (
	"github.com/jinzhu/gorm"
	"go-wsbackend/pkg/common"
)

var (
	cf *common.Config
	db *gorm.DB
)
func Init(c *common.Config) {
	cf = c
	db = cf.DB
}
