package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-wsbackend/pkg/common"
	"go-wsbackend/pkg/database"
	"go-wsbackend/pkg/model"
	"go-wsbackend/pkg/util"
	"go-wsbackend/spider/_36kr"
	"log"
	"strconv"
	"strings"
)

var (
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

func _36krOverview() {
	var infos []model.Information
	var dbInfos []model.Information
	_36krResp := _36kr.RunInformationOverview()
	if _36krResp.Code != 0 {
		log.Fatalf("fetch 36kr error: %d", _36krResp.Code)
	}

	db.Select("origin_id").Find(&dbInfos)
	dbOriginIDs := make([]string, len(dbInfos))
	for _, v := range dbInfos {
		dbOriginIDs  = append(dbOriginIDs, v.OriginID)
	}

	for _, item := range _36krResp.Data.Items{
		originID := strconv.FormatInt(item.ID, 10)
		if strings.Contains(item.Title, "垃圾") && util.IndexOf(dbOriginIDs, originID) == -1 {
			log.Printf("fetch waste information --> %s\n", item.Title)
			infos = append(infos, model.Information{
				OriginID: originID,
				Title: item.Title,
				SubTitle: item.Highlight.Content[0],
				Cover: item.TemplateInfo.TemplateCover[0],
			})
		}
	}
	if len(infos) > 0 {
		ids := make([]interface{}, len(infos))
		for i, v := range infos {
			ids[i] = v
		}
		_, err := database.BatchInsert(db, ids)
		if err != nil {
			panic(err)
		}
		fmt.Printf("save %d , successed! \n",len(ids))
	} else {
		fmt.Println("Nothing need add or update")
	}
}

func main(){
	_36krOverview()
	//fmt.Println(util.IndexOf([]string{"aaa", "a"}, "a"))
}
