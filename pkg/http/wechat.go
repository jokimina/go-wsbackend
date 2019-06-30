package http

import (
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/template"
	m "go-wsbackend/pkg/model"
	"log"
	"net/http"
)

func wechatValidate(c *gin.Context) {
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
	code := c.DefaultQuery("code", "")
	if code == "" {
		c.JSON(http.StatusBadRequest, m.ErrResponse{
			Status:  http.StatusBadRequest,
			Message: "",
		})
		return
	}
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

	message := &template.Message{
		ToUser:     openID,
		TemplateID: "c3M4soqvdhNZQQU0zHEWV2UIuLDjplKXmXd9XlzV850",
		FormID: formID,
		//URL: "",
		Data: map[string]*template.DataItem{
			"keyword1": {Value: "x"},
			"keyword2": {Value: "xx"},
			"keyword3": {Value: "xxx"},
		},
	}
	msgId, err := tpl.Send(message, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, m.ErrResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.String(http.StatusOK, string(msgId))
}
