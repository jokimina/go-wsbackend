package model

import "github.com/jinzhu/gorm"

type Information struct {
	gorm.Model
	Title string
	SubTitle string
	Content string `gorm:"type:text"`
	Cover string
	OriginID string `gorm:"type:varchar(100);unique_index"`
}

type InformationVo struct {
	Title string `json:"title"`
	SubTitle string `json:"sub_title"`
	Cover string `json:"cover"`
	OriginID string `json:"origin_id"`
}

