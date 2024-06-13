package main

import (
	"net/http"

	"example.com/shoppingCart-api/db"
	"example.com/shoppingCart-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/ticketLists", getTicketLists)

	server.Run(":8080")
}

func getTicketLists(context *gin.Context) {
	ticketLists, err := models.GetAllTicketLists()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch ticket lists", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, ticketLists)
}
