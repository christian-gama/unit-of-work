package main

import "github.com/gin-gonic/gin"

func init() {
	gin.SetMode(gin.ReleaseMode)
	print("\033[H\033[2J")
}

func main() {
	Bootstrap(gin.Default())
}
