package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-wsbackend/pkg/common"
	"go-wsbackend/pkg/model"
)

var conf common.Config

func Init(cf *common.Config){
	db, err := gorm.Open("mysql", "wsbackend:xiaodong@123@tcp(cashbustest.mysql.rds.aliyuncs.com)/wsbackend?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	cf.DB = db

	db.AutoMigrate(model.WasteItem{})
}
