package main

import (
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"learning/routes"
	"learning/utils/fireutil"
)

func main() {
	defer func(FirebaseClient *firestore.Client) {
		err := FirebaseClient.Close()
		if err != nil {
			panic(err)
		}
	}(fireutil.FirebaseClient)

	router := gin.Default()

	router.GET("/", routes.HomeRoute)
	router.GET("/transactions", routes.GetTransactions)
	router.POST("/transactions", routes.AddTransaction)
	router.DELETE("/transactions/:id", routes.DeleteTransaction)

	err := router.Run(":3000")
	if err != nil {
		panic(err)
	}
}
