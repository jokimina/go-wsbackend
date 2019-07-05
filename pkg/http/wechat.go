package http

import (
	"github.com/gin-gonic/gin"
	m "go-wsbackend/pkg/model"
	"go-wsbackend/pkg/service"
	"log"
	"net/http"
)

func wechatPush(c *gin.Context) {
	server := wechatSrv.GetServer(c.Request, c.Writer)
	err := server.Serve()
	if err != nil {
		c.JSON(http.StatusInternalServerError, m.ErrResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	server.Send()
	return
	//c.JSON(http.StatusOK, m.Response{})
}

func code2session(c *gin.Context) {
	appID := c.DefaultQuery("appid", "")
	code := c.DefaultQuery("code", "")
	if code == "" {
		c.JSON(http.StatusBadRequest, m.ErrResponse{
			Status:  http.StatusBadRequest,
			Message: "",
		})
		return
	}
	wxa := apps[appID].GetMiniProgram()
	r, err := wxa.Code2Session(code)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, m.ErrResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, m.Response{Status: http.StatusOK, Data: r.OpenID})
}

func sendAuditTemplate(c *gin.Context) {
	appID := c.DefaultQuery("appid", "")
	openID := c.DefaultQuery("openid", "")
	formID := c.DefaultQuery("formid", "")
	if openID == "" {
		c.JSON(http.StatusBadRequest, m.ErrResponse{
			Status:  http.StatusBadRequest,
			Message: "need openid!",
		})
		return
	}
	//r, err := wxa.Code2Session(code)

	tpl := apps[appID].GetTemplate()
	msgID, err := service.SendWechatTemplateMessage(tpl, &m.FeedbackBindObj{
		AppID:  appID,
		FormID: formID,
		OpenID: openID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, m.ErrResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.String(http.StatusOK, string(msgID))
}

func saveUserInfo(c *gin.Context) {
	var json m.BaseUserInfo
	var userInfo m.UserInfo
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, m.ErrResponse{Status:http.StatusBadRequest, Message: err.Error()})
		return
	}
	userInfo.BaseUserInfo = json
	if err := db.Save(&userInfo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, m.ErrResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, m.Response{Status: http.StatusOK})
}
