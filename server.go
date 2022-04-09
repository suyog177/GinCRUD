package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", GetString)

	r.Run()
}

func GetString(c *gin.Context) {

	c.String(200, "Hello world")
}
