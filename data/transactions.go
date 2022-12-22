package data

import (
	"context"
	"learning/models"
	"learning/utils/fireutil"
)

func DeleteTransactionById(id string) error {
	_, err := fireutil.FirebaseClient.Collection("transactions").Doc(id).Delete(context.Background())
	return err
}

func GetTransactionById(id string) *models.Transaction {
	snapshot, err := fireutil.FirebaseClient.Collection("transactions").Doc(id).Get(context.Background())
	if err != nil {
		return nil
	}
	return fireutil.FromSnapShotToStruct[models.Transaction](snapshot)
}

func GetTransactions(filter models.TransactionFilter) []models.Transaction {
	query := fireutil.FirebaseClient.Collection("transactions").Query
	if filter.Category != "" {
		query = query.Where("Category", "==", filter.Category)
	}
	documents := query.Documents(context.Background())
	return fireutil.FromDocumentsToSlice[models.Transaction](documents)
}

func AddTransaction(transaction *models.Transaction) error {
	ref := fireutil.FirebaseClient.Collection("transactions").NewDoc()
	transaction.Id = ref.ID
	_, err := fireutil.FirebaseClient.Collection("transactions").Doc(ref.ID).Set(context.Background(), transaction)
	return err
}
