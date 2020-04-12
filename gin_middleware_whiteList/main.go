package main

import "github.com/gin-gonic/gin"

//IPAuthMiddleware set whiteList for IP
func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string{
			"127.0.0.1",
			"::1",
		}
		flag := false
		clientIP := c.ClientIP()
		for _, host := range ipList {
			if clientIP == host {
				flag = true
			}
		}
		if !flag {
			c.String(401, "%s,not in iplist", clientIP)
			c.Abort()
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(IPAuthMiddleware())
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "%s", "hello test")
	})
	r.Run()
}
