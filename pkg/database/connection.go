package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-wsbackend/pkg/common"
	"go-wsbackend/pkg/model"
)

func Init(cf *common.Config) {
	m := cf.Mysql
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Url, m.DataBaseName)
	db, err := gorm.Open("mysql", s)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	cf.DB = db

	db.AutoMigrate(model.WasteItem{}, model.UserInfo{}, model.WasteSearchLog{})
}
