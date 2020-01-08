package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `form:"name"`
	Addr     string    `form:"addr"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", testing)
	r.POST("/testing", testing)
	r.Run()
}

func testing(c *gin.Context) {
	var person Person
	// content-type
	if err := c.ShouldBind(&person); err != nil {
		c.String(200, "person bind error: %v", err)
		c.Abort()
	}
	c.String(200, "%v", person)
}
