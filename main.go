package main

import (
	"log"
	"net/http"

	"github.com/DEMYSTIF/gin-postgres-api/controllers"
	"github.com/DEMYSTIF/gin-postgres-api/db"
	"github.com/DEMYSTIF/gin-postgres-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Fatal("Failed to connect the database")
	}

	db.AutoMigrate(&models.Certificate{})

	router := gin.Default()
	router.POST("/create", func(ctx *gin.Context) {
		controllers.CreateOne(ctx, db)
	})
	router.GET("/read", func(ctx *gin.Context) {
		controllers.ReadAll(ctx, db)
	})
	router.GET("/read/:id", func(ctx *gin.Context) {
		controllers.ReadOne(ctx, db)
	})
	router.PUT("/update/:id", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "Updated a certificate")
	})
	router.DELETE("/delete/:id", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "Deleted a certificate")
	})
	router.Run("localhost:8080")
}
