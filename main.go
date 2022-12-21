package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Transaction struct {
	Id       int    `json:"id"`
	Credit   int    `json:"credit"`
	Debit    int    `json:"debit"`
	Category string `json:"category"`
	Summary  string `json:"summary"`
}

type AddTransactionRequest struct {
	Credit   int    `json:"credit"`
	Debit    int    `json:"debit"`
	Category string `json:"category"`
	Summary  string `json:"summary"`
}

type ApiResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var transactions = []Transaction{
	{Id: 1, Credit: 2000, Debit: 0, Category: "Salary", Summary: "SamCom"},
	{Id: 2, Credit: 0, Debit: 350, Category: "Food", Summary: "Dinner"},
}

func homeRoute(context *gin.Context) {
	context.String(http.StatusOK, "Welcome to Go Modular")
}

func getTransactions(context *gin.Context) {
	var response ApiResponse
	if len(transactions) <= 0 {
		response = ApiResponse{
			Message: "Transactions not found",
			Data:    transactions,
		}
	} else {
		response = ApiResponse{
			Message: "Transactions found",
			Data:    transactions,
		}
	}

	context.JSON(http.StatusOK, response)
}

func addTransaction(context *gin.Context) {
	var addTransactionRequest AddTransactionRequest

	if err := context.BindJSON(&addTransactionRequest); err != nil {
		context.String(http.StatusInternalServerError, "Error parsing request")
		return
	}

	transaction := Transaction{
		Id:       int(time.Now().Unix()),
		Credit:   addTransactionRequest.Credit,
		Debit:    addTransactionRequest.Debit,
		Category: addTransactionRequest.Category,
		Summary:  addTransactionRequest.Summary,
	}
	transactions = append(transactions, transaction)

	response := ApiResponse{
		Message: "Transaction added",
		Data:    transaction,
	}

	context.JSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()

	router.GET("/", homeRoute)
	router.GET("/transactions", getTransactions)
	router.POST("/transactions", addTransaction)

	err := router.Run(":3000")
	if err != nil {
		panic(err)
	}
}
