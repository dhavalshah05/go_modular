package routes

import (
	"github.com/gin-gonic/gin"
	"learning/data"
	"learning/models"
	"net/http"
)

func GetTransactions(context *gin.Context) {
	var response models.ApiResponse
	filterCategory := context.Query("category")

	transactionFilter := models.TransactionFilter{
		Category: filterCategory,
	}
	transactions := data.GetTransactions(transactionFilter)

	if len(transactions) <= 0 {
		response = models.ApiResponse{
			Message: "Transactions not found",
			Data:    transactions,
		}
	} else {
		response = models.ApiResponse{
			Message: "Transactions found",
			Data:    transactions,
		}
	}

	context.JSON(http.StatusOK, response)
}
