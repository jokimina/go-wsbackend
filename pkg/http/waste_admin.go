package http

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
	m "go-wsbackend/pkg/model"
	"net/http"
	"strconv"
)

func fetchSingleWaste(c *gin.Context) {
	var waste []m.WasteItem
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id"},
		ShowSQL: true,
	}, &waste)
	c.JSON(http.StatusOK, paginator)
}
