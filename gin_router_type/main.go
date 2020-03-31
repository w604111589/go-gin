package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.String(200, "GET")
	})
	r.POST("/post", func(c *gin.Context) {
		c.String(200, "POST")
	})
	r.Run(":8082")
}
