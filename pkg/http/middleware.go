package http

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-wsbackend/pkg/common"
	"go-wsbackend/pkg/crypto"
	m "go-wsbackend/pkg/model"
	"net/http"
)

func respondWithError(message string, c *gin.Context) {

	c.JSON(http.StatusInternalServerError, m.ErrResponse{Status:http.StatusInternalServerError, Message: message})
	c.Abort()
}

func decryptMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params gin.Params
		paramMap := make(map[string]string)
		enc := c.DefaultQuery("enc", "")
		if enc != "" {
			src, err := base64.StdEncoding.DecodeString(enc)
			if err != nil {
				respondWithError(err.Error(), c)
			}
			b, err := crypto.DesDecrypt(src, common.Key)
			if err != nil {
				respondWithError(err.Error(), c)
			}
			json.Unmarshal(b, &paramMap)
			for k,v := range paramMap {
				params = append(params, gin.Param{
					Key: k,
					Value: v,
				})
			}
			c.Params = params
		}
		c.Next()
	}
}
