package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"go-wsbackend/pkg/common"
	"go-wsbackend/pkg/database"
	m "go-wsbackend/pkg/model"
	"io/ioutil"
	"os"
	"path"
)

type WasteData struct {
	Version int64         `json:"version"`
	Data    []m.WasteItem `json:"data"`
}

var (
	dataFile  string
	wasteData WasteData
	conf      *common.Config
	db        *gorm.DB
)

func init() {
	flag.StringVar(&dataFile, "f", path.Join("../..", "data", "data.json"), "")
	conf = &common.Config{
		DataFile: dataFile,
	}
	database.Init(conf)
	db = conf.DB
	db.Debug()
	db.LogMode(true)
}

func fileToDb() {
	jsonFile, err := os.Open(dataFile)
	defer jsonFile.Close()
	defer db.Close()
	if err != nil {
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &wasteData)

	var dbItems []m.WasteItem
	db.Select("id, name, cats", ).Find(&dbItems)

	var ds []m.WasteItem
	for _, data := range wasteData.Data {
		var sameItem m.WasteItem
		//db.Where(m.WasteItem{Name: data.Name}).First(&item)
		for _, dItem := range dbItems{
			if data.Name == dItem.Name{
				sameItem = dItem
				break
			}
		}
		if sameItem.ID != 0 {
			continue
		}
		ds = append(ds, data)
		fmt.Printf("Add item %s, cats: %d. \n", data.Name, data.Cats)
	}
	if len(ds) > 0 {
		ids := make([]interface{}, len(ds))
		for i, v := range ds {
			ids[i] = v
		}
		_, err = database.BatchInsert(db, ids)
		if err != nil {
			panic(err)
		}
		fmt.Printf("save version %v, count %d success \n", wasteData.Version, len(wasteData.Data))
	} else {
		fmt.Println("Nothing need add or update")
	}
	fmt.Println("Done.")

}

//fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))

func main() {
	fileToDb()
}
