package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/user/*aciton", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.Run(":8082")
}
