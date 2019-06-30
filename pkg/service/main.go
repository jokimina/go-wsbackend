package service

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"go-wsbackend/pkg/common"
	m "go-wsbackend/pkg/model"
	"io/ioutil"
	"log"
	"os"
)

var (
	cf           *common.Config
	db           *gorm.DB
	allWasteData *[]byte
	ws           []m.WasteItem
	dataJson     m.DataJson
)

func Init(c *common.Config) {
	cf = c
	db = cf.DB
	LoadAllDbWaste()
}

func LoadAllDbWaste() {
	log.Println("--> load database data...")
	if cf.UseMysql {
		db.Where(&m.WasteItem{Status: m.StatusOnline}).Find(&ws)
		dataJson.Version = 2
		dataJson.Data = ws
		b, err := json.Marshal(dataJson)
		if err != nil {
			panic(err)
		}
		allWasteData = &b
	} else {
		jsonFile, err := os.Open(cf.DataFile)
		defer jsonFile.Close()
		if err != nil {
			panic(err)
		}
		b, _ := ioutil.ReadAll(jsonFile)
		allWasteData = &b
	}
}
