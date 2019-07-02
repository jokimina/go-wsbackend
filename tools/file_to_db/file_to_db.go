package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"go-wsbackend/pkg/common"
	"go-wsbackend/pkg/database"
	m "go-wsbackend/pkg/model"
	"go-wsbackend/pkg/util"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
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
	url := util.DefaultGetEnvString("DB_URL", "127.0.0.1")
	username := util.DefaultGetEnvString("DB_USERNAME", "root")
	password := util.DefaultGetEnvString("DB_PASSWORD", "")
	dbName := util.DefaultGetEnvString("DB_NAME", "wsbackend")
	conf = &common.Config{
		DataFile: dataFile,
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

func fileToDbOfficial() {
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
		for _, dItem := range dbItems {
			if data.Name == dItem.Name {
				sameItem = dItem
				break
			}
		}
		if sameItem.ID != 0 {
			continue
		}
		data.From = m.FromOfficial
		data.Status = m.StatusOnline
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

func fileToDb3Data() {
	dataFile = path.Join("../..", "data", "3data.json")
	jsonFile, err := os.Open(dataFile)
	defer jsonFile.Close()
	defer db.Close()
	if err != nil {
		panic(err)
	}

	var json3Data m.Json3Data
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &json3Data)

	var all3Data []m.WasteItemVo
	all3Data = append(all3Data, json3Data.Num1...)
	all3Data = append(all3Data, json3Data.Num2...)
	all3Data = append(all3Data, json3Data.Num3...)
	all3Data = append(all3Data, json3Data.Num4...)

	var dbItems []m.WasteItem
	db.Select("id, name, cats", ).Find(&dbItems)

	// 不包含的结果
	var ds []m.WasteItem
	var dsNames []string
	var catMapping = map[int]int{
		4: 1,
		3: 2,
		2: 4,
		1: 3,
	}
	for _, data := range all3Data {
		var sameItem m.WasteItem
		//db.Where(m.WasteItem{Name: data.Name}).First(&item)
		for _, dItem := range dbItems {
			if strings.ToLower(data.N) == strings.ToLower(dItem.Name) {
				sameItem = dItem
				break
			}
		}

		if sameItem.ID != 0 || util.IndexOf(dsNames, data.N) != -1 {
			continue
		}
		cats, _ := strconv.Atoi(data.C)
		sameItem.Name = data.N
		sameItem.Cats = catMapping[cats]
		sameItem.From = m.FromWeApp
		sameItem.Status = m.StatusOnline
		qp, sp := util.GetPinYin(sameItem.Name)
		sameItem.Qp = qp
		sameItem.FL = sp
		ds = append(ds, sameItem)
		dsNames = append(dsNames, sameItem.Name)
		fmt.Printf("Add item %s, cats: %d. len: %d\n", sameItem.Name, sameItem.Cats, len(ds))
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

func main() {
	//fileToDbOfficial()
	//fileToDb3Data()
	type Test struct {
		Name string
		Test string
	}

	type Test2 struct {
		Test
	}

}
