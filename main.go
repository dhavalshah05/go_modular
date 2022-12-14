package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"learning/utils/boarding"
	"learning/utils/sheet"
)

func main() {
	err, year, month := boarding.GetYearAndMonthFromUser()
	handleError(err)

	file, err := excelize.OpenFile("BalanceSheet.xlsx")
	defer func() {
		// Close the spreadsheet.
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	handleError(err)

	categories := sheet.GetAllCategories(file)
	transactions := sheet.GetAllTransactions(file, year, month)

	if len(transactions) <= 0 {
		fmt.Printf("Cannot find transactions for Month: %d and Year: %d\n", month, year)
		return
	}

	var totalExpense = sheet.GetTotalExpense(transactions)
	for _, category := range categories {
		total := sheet.FindTotalExpenseInPercentageByCategory(category, transactions, totalExpense)
		fmt.Printf("%s : %f\n", category.Name, total)
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
