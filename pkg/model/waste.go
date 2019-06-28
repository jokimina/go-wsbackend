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
	// 可回收
	RecycleableWaste
	// 湿垃圾
	HouseholdWaste
	// 干垃圾
	ResidualWaste
)

var (
	// 三方小程序数据
	FromWeApp = "weapp"
	// 官方数据
	FromOfficial = "official"
	// 用户
	FromUser = "user"
	// 管理员手动后台添加
	FromAdmin = "admin"
	// 用户自定义提交的
	FromFeedback = "user"
)

type DataJson struct {
	Version uint8       `json:"version"`
	Data    []WasteItem `json:"data"`
}

// 垃圾信息主表
type WasteItem struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);unique_index;not null" json:"name"`
	Qp   string `json:"qp"`   // 全拼
	FL   string `json:"fl"`   // 首拼
	Cats int64  `json:"cats"` // 分类
	From string `json:"from"` // 数据来源
}

// 用户自己提交的记录
type UserCommitRecord struct {
	gorm.Model

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
