package model
//
//import (
//	"fmt"
//	"github.com/jinzhu/gorm"
//	"strings"
//)
//
//type Base struct {
//}
//
//func BulkInsert(db gorm.DB, unsavedRows []*interface{}) {
//	valueStrings := make([]string, 0, len(unsavedRows))
//	valueArgs := make([]interface{}, 0, len(unsavedRows)*3)
//	for _, post := range unsavedRows {
//		valueStrings = append(valueStrings, fmt.Sprintf("(%s)", strings.Repeat(",", len(unsavedRows))))
//		valueArgs = append(valueArgs, post.Column1)
//		valueArgs = append(valueArgs, post.Column2)
//		valueArgs = append(valueArgs, post.Column3)
//	}
//	stmt := fmt.Sprintf("INSERT INTO my_sample_table (column1, column2, column3) VALUES %s", strings.Join(valueStrings, ","))
//	fmt.Println(stmt)
//	//db.Exec(stmt, valueArgs...)
//}
//