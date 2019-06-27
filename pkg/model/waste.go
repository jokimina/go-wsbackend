package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-wsbackend/pkg/util"
	"strings"
)

type WasteItem struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);unique_index;not null" json:"name"`
	Qp   string `json:"qp"`
	FL   string `json:"fl"`
	Cats int64  `json:"cats"`
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
