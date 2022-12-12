package sheet

import (
	"github.com/xuri/excelize/v2"
	"strconv"
)

type Transaction struct {
	Date     string
	Credit   int
	Debit    int
	Category string
	Summary  string
}

type Category struct {
	Name string
}

func GetAllCategories(file *excelize.File) []Category {
	cols, err := file.GetCols("Database")
	handleError(err)

	var categories []Category

	for _, row := range cols {
		if row[0] == "Categories" {
			for index, col := range row {
				if index == 0 || col == "" {
					continue
				}
				categories = append(categories, Category{
					Name: col,
				})
			}
		}
	}
	return categories
}

func GetAllTransactions(file *excelize.File) []Transaction {
	rows, err := file.GetRows("Transactions")
	handleError(err)

	var transactions []Transaction

	for index, cols := range rows {
		if index == 0 {
			continue
		}
		date := getItemAtIndex(cols, 0)

		creditStr := getItemAtIndex(cols, 1)
		credit, err := strconv.Atoi(creditStr)
		if err != nil {
			credit = 0
		}

		debitStr := getItemAtIndex(cols, 2)
		debit, err := strconv.Atoi(debitStr)
		if err != nil {
			debit = 0
		}

		category := getItemAtIndex(cols, 3)
		summary := getItemAtIndex(cols, 4)

		transactions = append(transactions, Transaction{
			Date:     date,
			Credit:   credit,
			Debit:    debit,
			Category: category,
			Summary:  summary,
		})
	}
	return transactions
}

func getItemAtIndex(row []string, index int) string {
	var result string
	if len(row) >= index+1 {
		result = row[index]
	} else {
		result = ""
	}
	return result
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
