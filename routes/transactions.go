package routes

import (
	"github.com/gin-gonic/gin"
	"learning/data"
	"learning/models"
	"net/http"
	"time"
)

func GetTransactions(context *gin.Context) {
	var response models.ApiResponse
	transactions := data.GetTransactions()

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

func AddTransaction(context *gin.Context) {
	var addTransactionRequest models.AddTransactionRequest

	if err := context.BindJSON(&addTransactionRequest); err != nil {
		context.String(http.StatusInternalServerError, "Error parsing request")
		return
	}

	transaction := models.Transaction{
		Id:       int(time.Now().Unix()),
		Credit:   addTransactionRequest.Credit,
		Debit:    addTransactionRequest.Debit,
		Category: addTransactionRequest.Category,
		Summary:  addTransactionRequest.Summary,
	}
	data.AddTransaction(transaction)

	response := models.ApiResponse{
		Message: "Transaction added",
		Data:    transaction,
	}

	context.JSON(http.StatusOK, response)
}
