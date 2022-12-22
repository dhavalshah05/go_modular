package routes

import (
	"github.com/gin-gonic/gin"
	"learning/data"
	"learning/models"
	"net/http"
)

func DeleteTransaction(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, models.ErrorApiResponse("Provide transaction id"))
		return
	}

	transaction := data.GetTransactionById(id)
	if transaction == nil {
		context.JSON(http.StatusNotFound, models.ErrorApiResponse("Transaction not found for given id"))
		return
	}

	err := data.DeleteTransactionById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, models.ErrorApiResponse("Error while deleting transaction"))
		return
	}

	response := models.ApiResponse{
		Message: "Transaction deleted",
		Data:    nil,
	}

	context.JSON(http.StatusOK, response)
}
