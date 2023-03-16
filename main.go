package main

import "github.com/gin-gonic/gin"

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	Bootstrap(gin.Default())
}
