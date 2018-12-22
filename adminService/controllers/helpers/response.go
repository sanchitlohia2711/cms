package helpers

import (
	"github.com/gin-gonic/gin"
)

//WriteResponse writes the response back
func WriteResponse(c *gin.Context, body interface{}, httpCode int) {
	c.JSON(httpCode, body)
}
