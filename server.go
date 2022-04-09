package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", GetString)
	r.GET("/:name", GetName)

	r.Run()
}

func GetString(c *gin.Context) {

	c.String(200, "Hello world")
}

func GetName(c *gin.Context) {
	name := c.Param("name")
	msg := fmt.Sprintf("Hello world, %s", name)
	c.String(200, msg)
}
