package util

import (
	"github.com/mozillazg/go-pinyin"
	"strings"
)

func GetPinYin(s string)(string, string){
	var rqp []string
	// 首拼
	var rsp []string
	a := pinyin.NewArgs()
	ps := pinyin.Pinyin(s, a)
	for _, v := range ps {
		rqp = append(rqp, v[0])
		rsp = append(rsp, strings.Split(v[0], "")[0])
	}
	return strings.Join(rqp, ""), strings.Join(rsp, "")
}

