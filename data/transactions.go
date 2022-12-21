package data

import (
	"learning/models"
)

var storedTransactions = []models.Transaction{
	{Id: 1, Credit: 2000, Debit: 0, Category: "Salary", Summary: "SamCom"},
	{Id: 2, Credit: 0, Debit: 350, Category: "Food", Summary: "Dinner"},
}

func GetTransactions() []models.Transaction {
	return storedTransactions
}

func AddTransaction(transaction models.Transaction) {
	storedTransactions = append(storedTransactions, transaction)
}
