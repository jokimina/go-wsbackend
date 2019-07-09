package service

import (
	m "go-wsbackend/pkg/model"
	"strings"
)

func GetEncData() []byte {
	return encData
}

func GetWasteCount() uint16 {
	return wasteCount
}

func Search(s string) []m.WasteShortVo {
	r := make([]m.WasteShortVo, 0)
	for _, item := range wasteItemVoList {
		if strings.Index(item.Name, s) > -1 || strings.Index(item.Qp, s) > -1 || strings.Index(item.FL, s) > -1 {
			r = append(r, m.WasteShortVo{
				Name: item.Name,
				Cats: item.Cats,
				From: item.From,
			})
		}
	}
	return r
}
