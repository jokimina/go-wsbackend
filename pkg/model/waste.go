package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-wsbackend/pkg/util"
	"strings"
)

type WasteSorting int

const (
	// 有害垃圾
	HazardousWaste = iota + 1
	// 可回收物
	RecycleableWaste
	// 湿垃圾
	HouseholdWaste
	// 干垃圾
	ResidualWaste
)

var Wastes = []string{"有害垃圾","可回收物", "湿垃圾", "干垃圾"}

func GetWasteNameByIndex(i int) string{
	return Wastes[i - 1]
}

var (
	// 三方小程序数据
	FromWeApp = "weapp"
	// 官方数据
	FromOfficial = "official"
	// 用户
	FromUser = "user"
	// 管理员手动后台添加
	FromAdmin = "admin"

	// Status
	// 显示到线上的版本
	StatusOnline = "online"
	// 下线
	StatusOffline = "offline"
	// 审核中
	StatusPendding = "pending"
	// 取消
	StatusCancel = "cancel"
)

type DataJson struct {
	Version uint8       `json:"version"`
	Data    []WasteItem `json:"data"`
}

// 垃圾信息主表
type WasteItem struct {
	gorm.Model
	Name   string `gorm:"type:varchar(100);unique_index;not null" json:"name"`
	Qp     string `json:"qp"`   // 全拼
	FL     string `json:"fl"`   // 首拼
	Cats   int  `json:"cats"` // 分类
	From   string `json:"-"` // 数据来源
	FormID   string `json:"-"` // 小程序 form_id
	OpenID   string `json:"-"` // 小程序 open_id
	AppID   string `json:"-"` // 小程序 appid
	Status string `gorm:"default:'online'" json:"status"'`
}

type WasteItemVo struct {
	N string `json:"n"` //名称
	Q string `json:"q"` // 全拼
	F string `json:"f"` // 首拼
	C string `json:"c"` // 分类
}

type FeedbackBindObj struct {
	Name   string `json:"name"`
	Cats   int  `json:"cats"`
	OpenID string `json:"open_id"`
	FormID string `json:"form_id"`
	AppID string `json:"appid"`
}

func (m WasteItem) BulkInsert(db *gorm.DB, ws []WasteItem) error {
	tableName := db.NewScope(m).GetModelStruct().TableName(db)
	fieldStr := ""
	fields := util.GetFields(&m, true)
	fieldNums := len(fields)
	quesMarkString := "("
	for i := 0; i < fieldNums; i++ {
		if i != (fieldNums - 1) {
			fieldStr += gorm.ToColumnName(fields[i].Field.Name) + ","
		} else {
			fieldStr += gorm.ToColumnName(fields[i].Field.Name)
		}
		quesMarkString += "?,"
	}
	quesMarkString = quesMarkString[:len(quesMarkString)-1] + ")"

	valueStrings := make([]string, 0, len(ws))
	valueArgs := make([]interface{}, 0, len(ws)*fieldNums)

	for _, w := range ws {
		valueStrings = append(valueStrings, quesMarkString)
		for _, field := range fields {
			f := util.GetFieldByName(w, field.Field.Name)
			valueArgs = append(valueArgs, f.Interface())
		}
	}

	stmt := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s", tableName, fieldStr, strings.Join(valueStrings, ","))
	fmt.Println(stmt)
	fmt.Println(valueArgs...)
	err := db.Exec(stmt, valueArgs...).Error
	return err
}
