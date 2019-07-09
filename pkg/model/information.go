package model

import "github.com/jinzhu/gorm"

type Information struct {
	gorm.Model
	Title string
	SubTitle string
	Content string
	Cover string
	OriginID string `gorm:"type:varchar(100);unique_index"`
}

type InformationDetail struct {

}