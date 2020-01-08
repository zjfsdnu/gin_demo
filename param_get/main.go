package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		firstName := c.Query("first_name")
		lastName := c.DefaultQuery("last_name", "last_default_name")
		c.String(200, "%s, %s", firstName, lastName)
	})
	r.Run()
}
