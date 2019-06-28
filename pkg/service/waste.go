package service

import (
	"encoding/base64"
	"encoding/json"
	"go-wsbackend/pkg/common"
	"go-wsbackend/pkg/crypto"
	m "go-wsbackend/pkg/model"
	"io/ioutil"
	"os"
)

func GetEncData(dataFile string) (encData []byte) {
	var (
		bt *[]byte
		ws []m.WasteItem
		dataJson m.DataJson
	)

	if cf.UseMysql {
		db.Find(&ws)
		dataJson.Version = 2
		dataJson.Data = ws
		b, err := json.Marshal(dataJson)
		if err != nil {
			panic(err)
		}
		bt = &b
	} else {
		jsonFile, err := os.Open(dataFile)
		defer jsonFile.Close()
		if err != nil {
			panic(err)
		}
		b, _ := ioutil.ReadAll(jsonFile)
		bt = &b
	}

	outs, _ := crypto.DesEncrypt(*bt, common.Key)
	encData = make([]byte, base64.StdEncoding.EncodedLen(len(outs)))
	base64.StdEncoding.Encode(encData, outs)
	return
}
