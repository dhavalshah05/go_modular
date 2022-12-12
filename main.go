package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"learning/utils/sheet"
)

func main() {
	file, err := excelize.OpenFile("BalanceSheet.xlsx")
	defer func() {
		// Close the spreadsheet.
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	handleError(err)

	categories := sheet.GetAllCategories(file)
	transactions := sheet.GetAllTransactions(file)

	transactions = filterByCategory(categories[1], transactions)
	printTransactions(transactions)

	total := findTotalExpenseByCategory(categories[1], transactions)
	fmt.Printf("Total: %d\n", total)

	/*for _, category := range categories {
		total := findTotalExpenseByCategory(category, transactions)
		fmt.Printf("%s : %d\n", category.Name, total)
	}*/
}

func filterByCategory(category sheet.Category, transactions []sheet.Transaction) []sheet.Transaction {
	var result []sheet.Transaction
	for _, transaction := range transactions {
		if transaction.Category == category.Name {
			result = append(result, transaction)
		}
	}
	return result
}

func findTotalExpenseByCategory(category sheet.Category, transactions []sheet.Transaction) int {
	filteredTransactions := filterByCategory(category, transactions)

	var result = 0
	for _, transaction := range filteredTransactions {
		result = result + transaction.Debit
	}
	return result
}

func printTransactions(transactions []sheet.Transaction) {
	for _, transaction := range transactions {
		fmt.Printf("%+v\n", transaction)
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
