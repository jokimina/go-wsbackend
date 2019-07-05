package http

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
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
