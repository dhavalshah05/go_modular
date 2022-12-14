package sheet

import (
	"github.com/xuri/excelize/v2"
	"strconv"
	"time"
)

type Transaction struct {
	Date     time.Time
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

func GetAllTransactions(file *excelize.File, year int, month int) []Transaction {
	rows, err := file.GetRows("Transactions")
	handleError(err)

	var transactions []Transaction

	for index, cols := range rows {
		if index == 0 {
			continue
		}
		date := getItemAtIndex(cols, 0)
		parsedDate, err := time.Parse("2-Jan-2006", date)
		if err != nil {
			panic(err)
		}

		if parsedDate.Year() != year && parsedDate.Month() != time.Month(month) {
			continue
		}

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
			Date:     parsedDate,
			Credit:   credit,
			Debit:    debit,
			Category: category,
			Summary:  summary,
		})
	}
	return transactions
}

func GetTotalExpense(transactions []Transaction) int {
	var total = 0
	for _, transaction := range transactions {
		total = total + transaction.Debit
	}
	return total
}

func GetTotalIncome(transactions []Transaction) int {
	var total = 0
	for _, transaction := range transactions {
		total = total + transaction.Credit
	}
	return total
}

func FilterByCategory(category Category, transactions []Transaction) []Transaction {
	var result []Transaction
	for _, transaction := range transactions {
		if transaction.Category == category.Name {
			result = append(result, transaction)
		}
	}
	return result
}

func FindTotalExpenseInRsByCategory(category Category, transactions []Transaction) int {
	filteredTransactions := FilterByCategory(category, transactions)

	var result = 0
	for _, transaction := range filteredTransactions {
		result = result + transaction.Debit
	}
	return result
}

func FindTotalExpenseInPercentageByCategory(category Category, transactions []Transaction, totalExpense int) float64 {
	var result = FindTotalExpenseInRsByCategory(category, transactions)
	return (float64(result) * 100.0) / float64(totalExpense)
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
