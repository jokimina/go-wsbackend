package http

import (
	"github.com/gin-gonic/gin"
	m "go-wsbackend/pkg/model"
	"net/http"
)

func getInformationList(c *gin.Context) {
	var infoList []m.InformationVo
	var dbInfoList []m.Information
	if err := db.Limit(20).Find(&dbInfoList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, m.ErrResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	for _, v := range dbInfoList {
		infoList = append(infoList, m.InformationVo{
			Title: v.Title,
			SubTitle: v.SubTitle,
			Cover: v.Cover,
			OriginID: v.OriginID,
		})
	}
	c.JSON(http.StatusOK, m.Response{Status:http.StatusOK, Data: infoList})
}

func getInformationContent(c *gin.Context) {
	var info m.Information
	originID := c.Param("originID")
	if originID != "" {
		if err := db.Where("origin_id = ?", originID).Find(&info).Error; err != nil {
			c.JSON(http.StatusInternalServerError, m.ErrResponse{Status: http.StatusInternalServerError, Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, m.Response{Status: http.StatusOK, Data: info.Content})
		return
	}
	c.JSON(http.StatusBadRequest, m.ErrResponse{Status:http.StatusBadRequest, Message: "bad param!"})
}
