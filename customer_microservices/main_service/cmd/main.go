package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// customer := r.Group("customer")

	r.Run()
}
