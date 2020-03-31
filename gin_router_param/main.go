package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		firstName := c.Query("firstName")
		lastName := c.DefaultQuery("lastName", "deault_value")
		c.String(http.StatusOK, "%s,%s", firstName, lastName)

	})
	r.POST("/test/post", func(c *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		// c.String(http.StatusOK, string(bodyBytes))
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		firstName := c.PostForm("firstName")
		lastName := c.DefaultPostForm("lastName", "default_value")
		c.String(http.StatusOK, "%s,%s,%s", firstName, lastName, string(bodyBytes))

	})
	r.Run(":8082")
}
