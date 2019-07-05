package model

import "github.com/jinzhu/gorm"

type BaseUserInfo struct {
	NickName  string `json:"nickName"`
	Gender    int64  `json:"gender"`
	Language  string `json:"language"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatarUrl"`
	Openid    string `gorm:"type:varchar(100);unique_index;not null" json:"openid"`
}

type UserInfo struct {
	gorm.Model
	BaseUserInfo
}

