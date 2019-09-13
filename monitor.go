package main

import (
	"math/rand"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// c.Writer.Write([]byte(fmt.Sprintf("%s", time.Now().String())))
	// c.Writer.Write([]byte(fmt.Sprintf("%d", r.Int63n(1000))))
	var m = make(map[string]interface{}, 1)
	m["qps"] = r.Int63n(1000)
	c.JSON(200, m)
	return
}

func main() {
	router := gin.Default()

	// router.PUT("/upload", Upload)
	router.GET("/upload", Upload)

	dir := "./build"
	router.Use(static.Serve("/", static.LocalFile(dir, true)))

	// curl xxx:9081/a
	// routers.Use(static.Serve("/a", static.LocalFile(*dir, true)))

	router.Run(":8083")
}
