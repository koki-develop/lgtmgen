package api

import "github.com/gin-gonic/gin"

func NewEngine() *gin.Engine {
	r := gin.Default()

	r.GET("/h", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	return r
}
