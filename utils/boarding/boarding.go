package boarding

import (
	"fmt"
	"learning/utils/input"
	"strconv"
)

func GetYearAndMonthFromUser() (err error, year int, month int) {
	fmt.Print("Enter year: ")
	yearStr := input.GetStringFromUser()
	fmt.Print("Enter month: ")
	monthStr := input.GetStringFromUser()

	year, err = strconv.Atoi(yearStr)
	if err != nil {
		return err, 0, 0
	}

	month, err = strconv.Atoi(monthStr)
	if err != nil {
		return err, 0, 0
	}

	return nil, year, month
}
