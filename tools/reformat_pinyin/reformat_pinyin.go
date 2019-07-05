package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-wsbackend/pkg/common"
	"go-wsbackend/pkg/database"
	"go-wsbackend/pkg/model"
	"go-wsbackend/pkg/util"
)

var (
	dataFile string
	conf     *common.Config
	db       *gorm.DB
)

func init() {
	url := util.DefaultGetEnvString("DB_URL", "127.0.0.1")
	username := util.DefaultGetEnvString("DB_USERNAME", "root")
	password := util.DefaultGetEnvString("DB_PASSWORD", "")
	dbName := util.DefaultGetEnvString("DB_NAME", "wsbackend")
	conf = &common.Config{
		UseMysql: true,
		Mysql: struct {
			Url          string
			Username     string
			Password     string
			DataBaseName string
		}{Url: url, Username: username, Password: password, DataBaseName: dbName},
	}
	database.Init(conf)
	db = conf.DB
	db.Debug()
	db.LogMode(true)
}

func main() {
	var wsAll []model.WasteItem
	//db.Limit(10).Find(&wsAll)
	db.Find(&wsAll)
	util.Transact(db, func(tx *gorm.DB) error {
		for _, item := range wsAll {
			fmt.Println(item)
			qp, sp := util.GetPinYin(item.Name)
			item.Qp = qp
			item.FL = sp
			if err := tx.Model(&item).Update("qp", "fl").Error; err != nil {
				return err
			}
		}
		return tx.Error
	})
}
