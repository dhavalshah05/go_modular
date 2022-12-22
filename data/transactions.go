package data

import (
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"learning/models"
	"learning/utils"
)

func GetTransactions(filter models.TransactionFilter) []models.Transaction {
	query := utils.FirebaseClient.Collection("transactions").Query
	if filter.Category != "" {
		query = query.Where("Category", "==", filter.Category)
	}
	documents := query.Documents(context.Background())
	var transactions []models.Transaction
	for {
		next, err := documents.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		var result models.Transaction
		err = next.DataTo(&result)
		if err == nil {
			transactions = append(transactions, result)
		}
	}
	return transactions
}

func AddTransaction(transaction *models.Transaction) error {
	ref := utils.FirebaseClient.Collection("transactions").NewDoc()
	transaction.Id = ref.ID
	_, err := utils.FirebaseClient.Collection("transactions").Doc(ref.ID).Set(context.Background(), transaction)
	return err
}
