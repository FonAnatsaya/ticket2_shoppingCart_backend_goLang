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
	server.POST("/discount", applyDiscount)

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

func applyDiscount(context *gin.Context) {

	var request struct {
		DiscountCode string  `json:"discountCode"`
		TotalPrice   float64 `json:"totalPrice"`
	}

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	discount, err := models.GetDiscountValue(request.DiscountCode, request.TotalPrice)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not apply discount", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"discount": discount})
}
