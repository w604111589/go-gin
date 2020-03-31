package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		//return c.JSON(200,gin.H{"code":200,"data":"","msg":"ok"})
		c.JSON(200, gin.H{"msg": "ping"})
	})
	r.Run()
}
