package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(IPAuthMiddleware())
	r.GET("test", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	r.Run()
}

func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipWhiteList := []string{
			"127.0.0.2",
		}
		flag := false
		clientIP := c.ClientIP()

		for _, host := range ipWhiteList {
			if clientIP == host {
				flag = true
				break
			}
		}
		if !flag {
			c.String(401, "%s, not in ipWhitelist", clientIP)
			c.Abort()
		}
	}
}
