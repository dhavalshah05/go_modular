package routes

import (
	"github.com/gin-gonic/gin"
	"learning/data"
	"learning/models"
	"net/http"
)

func AddTransaction(context *gin.Context) {
	var addTransactionRequest models.AddTransactionRequest

	if err := context.BindJSON(&addTransactionRequest); err != nil {
		context.JSON(http.StatusInternalServerError, models.ErrorApiResponse("Error parsing request"))
		return
	}

	if addTransactionRequest.Credit <= 0 && addTransactionRequest.Debit <= 0 {
		context.JSON(http.StatusBadRequest, models.ErrorApiResponse("credit or debit value is required"))
		return
	}

	if addTransactionRequest.Category == "" {
		context.JSON(http.StatusBadRequest, models.ErrorApiResponse("category is required"))
		return
	}

	if addTransactionRequest.Summary == "" {
		context.JSON(http.StatusBadRequest, models.ErrorApiResponse("summary is required"))
		return
	}

	transaction := models.Transaction{
		Id:       "",
		Credit:   addTransactionRequest.Credit,
		Debit:    addTransactionRequest.Debit,
		Category: addTransactionRequest.Category,
		Summary:  addTransactionRequest.Summary,
	}

	err := data.AddTransaction(&transaction)
	if err != nil {
		context.JSON(http.StatusInternalServerError, models.ErrorApiResponse("Unable to save transaction to firebase"))
		return
	}

	response := models.ApiResponse{
		Message: "Transaction added",
		Data:    transaction,
	}

	context.JSON(http.StatusOK, response)
}
