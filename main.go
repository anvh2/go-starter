package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pang")
	})
	// pprof.Register(r)
	return r
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	// i, _ := testErr()
	// println(i)

	// log.NewLog("test.log")
	// log.Log.Println("Running")
	r := setupRouter()
	r.Run(":8088")
}
