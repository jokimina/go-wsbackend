package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-wsbackend/pkg/service"
)

func getAllWaste(c *gin.Context){
	encData := service.GetEncData(cf.DataFile)
	c.String(http.StatusOK, string(encData))
}
