package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "go-wsbackend/pkg/model"
	"go-wsbackend/pkg/pagination"
	"go-wsbackend/pkg/service"
	"net/http"
	"strconv"
)

type auditReqJson struct{
	ID int `json:"ID"`
	Cats int `json: cats`
	Status string `json: status`
}

func fetchWaste(c *gin.Context) {
	var waste []m.WasteItem
	var cdb *gorm.DB
	var status = c.DefaultQuery("status", m.StatusOnline)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	name := c.DefaultQuery("name", "")
	cdb = db.Where("name like ? and status = ?", fmt.Sprintf("%%%s%%", name), status)

	paginator := pagination.Paging(&pagination.Param{
		DB:      cdb,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id"},
		ShowSQL: true,
	}, &waste)
	c.JSON(http.StatusOK, m.Response{
		Status: http.StatusOK,
		Data: paginator,
	})
}

func addWaste(c *gin.Context) {
	var waste m.WasteItem
	var rWaste m.WasteItem
	err := c.BindJSON(&waste)
	if err != nil {
		c.JSON(http.StatusInternalServerError, m.ErrResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	db.Where(&m.WasteItem{Name: waste.Name}).First(&rWaste)
	if rWaste.ID == 0 {
		waste.From = m.FromAdmin
		waste.Status = m.StatusOnline
		db.Create(&waste)
		c.JSON(http.StatusOK, m.Response{Status:http.StatusOK, Data:"saved"})
	} else {
		c.JSON(http.StatusConflict, m.ErrResponse{Status:http.StatusConflict, Message:"Already exists"})
	}
}

func updateWaste(c *gin.Context) {
	id := c.Param("id")
	var waste m.WasteItem
	db.First(&waste, id)
	if waste.ID == 0 {
		c.JSON(http.StatusNotFound, m.ErrResponse{Status: http.StatusNotFound, Message: "No waste found!"})
		return
	}
	err := c.BindJSON(&waste)
	if err != nil {
		c.JSON(http.StatusNotFound, m.ErrResponse{Status: http.StatusNotFound, Message: err.Error()})
		return
	}
	db.Save(&waste)
	c.JSON(http.StatusOK, m.Response{Status:http.StatusOK, Data:""})
}

func reloadWaste(c *gin.Context) {
	service.LoadAllDbWaste()
	c.JSON(http.StatusOK, m.Response{Status:http.StatusOK, Data:""})
}

func auditWaste(c *gin.Context) {
	var json auditReqJson
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var waste m.WasteItem
	db.First(&waste, json.ID)
	if waste.ID == 0 {
		c.JSON(http.StatusNotFound, m.ErrResponse{Status: http.StatusNotFound, Message: "No waste found!"})
		return
	}
	if json.Status != m.StatusOnline && json.Status != m.StatusDeny {
		c.JSON(http.StatusBadRequest, m.ErrResponse{Status: http.StatusBadRequest, Message: "param error!"})
		return
	}
	waste.Status = json.Status
	waste.Cats = json.Cats
	db.Save(&waste)
	c.JSON(http.StatusOK, m.Response{Status:http.StatusOK, Data:""})
}

func auditBatchWaste(c *gin.Context) {
	var json []auditReqJson
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, item := range json {
		var waste m.WasteItem
		db.First(&waste, item.ID)
		if waste.ID == 0 {
			c.JSON(http.StatusNotFound, m.ErrResponse{Status: http.StatusNotFound, Message: "No waste found!"})
			return
		}
		if item.Status != m.StatusOnline && item.Status != m.StatusDeny {
			c.JSON(http.StatusBadRequest, m.ErrResponse{Status: http.StatusBadRequest, Message: "param error!"})
			return
		}
		waste.Status = item.Status
		waste.Cats = item.Cats
		db.Save(&waste)
	}
	c.JSON(http.StatusOK, m.Response{Status:http.StatusOK, Data:""})
}