package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgres://demystif:gppw2023@localhost:5432/gin-postgres"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect the database")
	}
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