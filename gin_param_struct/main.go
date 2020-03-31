package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", testing)
	r.POST("/testing", testing)
	r.Run(":8082")
}

func testing(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err != nil {
		c.String(200, "person bind error", err)
	} else {
		c.String(200, "%v", person)
	}
}
