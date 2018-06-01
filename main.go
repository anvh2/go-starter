package main

import (
	log "github.com/bakaoh/go-common"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pang")
	})
	return r
}

func main() {
	log.NewLog("test.log")
	log.Log.Println("Running")
	r := setupRouter()
	r.Run(":8080")
}
