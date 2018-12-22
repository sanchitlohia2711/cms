package controllers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ko/cms-db/adminService/controllers/v1"
)

//SetUpRoutes set all the routes
func SetUpRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//Set v1 routes
	v1.SetUpRoutes(r)
	return
}
