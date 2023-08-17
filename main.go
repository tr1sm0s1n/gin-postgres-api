package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/create", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusCreated, "Created a certificate")
	})
	router.GET("/read", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "Read all certificates")
	})
	router.GET("/read/:id", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "Read a certificate")
	})
	router.PUT("/update/:id", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "Updated a certificate")
	})
	router.DELETE("/delete/:id", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "Deleted a certificate")
	})
	router.Run("localhost:8080")
}