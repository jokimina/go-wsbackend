package http

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-wsbackend/pkg/database"
	m "go-wsbackend/pkg/model"
	"go-wsbackend/pkg/service"
	"go-wsbackend/pkg/util"
	"net/http"
)

func getWasteChecksum(c *gin.Context) {
	encData := service.GetEncData()
	checksum := fmt.Sprintf("%x", md5.Sum(encData))
	c.String(http.StatusOK, checksum)
}

func getAllWaste(c *gin.Context) {
	encData := service.GetEncData()
	c.String(http.StatusOK, string(encData))
}

func searchWaste(c *gin.Context) {
	s:= c.Param("s")
	if s == "" {
		c.JSON(http.StatusBadRequest, m.ErrResponse{Status: http.StatusBadRequest, Message: "param error!"})
		return
	}
	go func(s string) {
		var wsl m.WasteSearchLog
		_ = database.Transact(db, func(tx *gorm.DB) error {
			tx.Where("s = ?", s).First(&wsl)
			if wsl.ID != 0 {
				tx.Model(&wsl).Update(&m.WasteSearchLog{S: s, C: wsl.C + 1})
			} else {
				tx.Save(&m.WasteSearchLog{S: s, C: 1})
			}
			return nil
		})
	}(s)
	searchResult := service.Search(s)
	// 限制返回数量 防止拖库
	if len(searchResult) > 100 {
		c.JSON(http.StatusInternalServerError, m.ErrResponse{Status: http.StatusInternalServerError, Message: "result too long!"})
		return
	}
	c.JSON(http.StatusOK, m.Response{Status: http.StatusOK, Data: searchResult})
}

func getWasteCount(c *gin.Context) {
	c.JSON(http.StatusOK, m.Response{Status:http.StatusOK, Data: service.GetWasteCount()})
}

func userFeedback(c *gin.Context) {
	var ws m.WasteItem
	var dbWs m.WasteItem
	var feedbackBindObj m.FeedbackBindObj
	if err := c.ShouldBindJSON(&feedbackBindObj); err != nil {
		c.JSON(http.StatusBadRequest, m.ErrResponse{Status:http.StatusBadRequest, Message: err.Error()})
		return
	}
	if db.Where(&m.WasteItem{Name: feedbackBindObj.Name}).First(&dbWs); dbWs.ID != 0 {
		c.JSON(http.StatusConflict, m.ErrResponse{Status:http.StatusConflict, Message:"Already exists"})
		return
	}
	qp, sp := util.GetPinYin(feedbackBindObj.Name)
	ws.Name = feedbackBindObj.Name
	ws.Cats = feedbackBindObj.Cats
	ws.FormID = feedbackBindObj.FormID
	ws.AppID = feedbackBindObj.AppID
	ws.OpenID = feedbackBindObj.OpenID
	ws.Status = m.StatusPending
	ws.From = m.FromUser
	ws.Qp = qp
	ws.FL = sp
	db.Create(&ws)
	//msgID, err := service.SendWechatTemplateMessage(tpl, &feedbackBindObj)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, m.ErrResponse{Status:http.StatusInternalServerError, Message: err.Error()})
	//	return
	//}
	c.JSON(http.StatusOK, m.Response{Status:http.StatusOK})
}
