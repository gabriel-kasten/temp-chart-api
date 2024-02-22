package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temp-chart-go/database"

	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	db, err := database.DatabaseCon()
	if err != nil {
		log.Fatal(err)
	}

	app := gin.Default()
	app.Use(cors.Default())

	app.GET("/temperature-data", func(ctx *gin.Context) {

		results, err := database.QueryTemperatureData(db)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, results)
	})
	app.Run(":6969")
}
