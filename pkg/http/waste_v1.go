package http

import (
	"github.com/gin-gonic/gin"
	"go-wsbackend/pkg/service"
	"net/http"
)

func getAllWaste(c *gin.Context) {
	encData := service.GetEncData()
	c.String(http.StatusOK, string(encData))
}
