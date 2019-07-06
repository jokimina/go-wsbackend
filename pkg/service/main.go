package service

import (
	"encoding/base64"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"go-wsbackend/pkg/common"
	"go-wsbackend/pkg/crypto"
	m "go-wsbackend/pkg/model"
	"io/ioutil"
	"log"
	"os"
)

var (
	cf           *common.Config
	db           *gorm.DB
	allWasteData *[]byte
	wasteItemVoList []m.WasteItemVo
	encData      []byte
	wasteCount	uint16
)

func Init(c *common.Config) {
	cf = c
	db = cf.DB
	LoadAllDbWaste()
}

func LoadAllDbWaste() {
	var (
		ws           []m.WasteItem
		dataJson     m.JsonData
	)
	log.Println("--> load database data...")
	if cf.UseMysql {
		db.Where(&m.WasteItem{Status: m.StatusOnline}).Find(&ws)
		for _, v := range ws {
			wasteItemVoList = append(wasteItemVoList, m.WasteItemVo{
				Name: v.Name,
				Qp: v.Qp,
				FL: v.FL,
				Cats: v.Cats,
				From: v.From,
			})
		}
		wasteCount = uint16(len(wasteItemVoList))
		dataJson.Version = 2
		dataJson.Data = wasteItemVoList
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
	outs, _ := crypto.DesEncrypt(*allWasteData, common.Key)
	encData = make([]byte, base64.StdEncoding.EncodedLen(len(outs)))
	base64.StdEncoding.Encode(encData, outs)
}
