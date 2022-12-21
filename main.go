package main

import (
	"github.com/gin-gonic/gin"
	"learning/routes"
)

func main() {
	router := gin.Default()

	router.GET("/", routes.HomeRoute)
	router.GET("/transactions", routes.GetTransactions)
	router.POST("/transactions", routes.AddTransaction)

	err := router.Run(":3000")
	if err != nil {
		panic(err)
	}
}
