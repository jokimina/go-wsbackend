package http

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-wsbackend/pkg/pagination"
	"github.com/gin-gonic/gin"
	m "go-wsbackend/pkg/model"
	"net/http"
	"strconv"
)

func fetchWaste(c *gin.Context) {
	var waste []m.WasteItem
	var cdb *gorm.DB
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	name := c.DefaultQuery("name", "")
	cdb = db.Where("name like ?", fmt.Sprintf("%%%s%%", name))

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