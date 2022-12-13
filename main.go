package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	nowUtc := time.Now().UTC()

	printTimeManually(now)
	printTimeManually(nowUtc)
}

func printTimeUsingFormat(datetime time.Time) {
	fmt.Println("Year :", datetime.Format("2006/06"))
	fmt.Println("Month :", datetime.Format("01/Jan/January"))
	fmt.Println("Day of month :", datetime.Format("02/Mon/Monday"))
	fmt.Println("Hour 24 :", datetime.Format("15"))
	fmt.Println("Hour 12 :", datetime.Format("03 PM/pm"))
	fmt.Println("Minutes :", datetime.Format("04"))
	fmt.Println("Seconds :", datetime.Format("05"))
	fmt.Println("Timezone :", datetime.Format("-07/-0700/MST"))
}

func printTimeUsingDate() {
	datetime := time.Date(2022, 12, 13, 8, 30, 0, 0, time.Local)
	printTimeManually(datetime)
}

func printTimeNow() {
	datetime := time.Now()
	printTimeManually(datetime)
}

func printTimeManually(datetime time.Time) {
	fmt.Println(datetime)
	fmt.Printf("Year: %d, Month: %d, Day: %d\n", datetime.Year(), datetime.Month(), datetime.Day())
	fmt.Printf("Hour: %d, Minute: %d, Second: %d\n", datetime.Hour(), datetime.Minute(), datetime.Second())
	fmt.Printf("Weekday: %v\n", datetime.Weekday())
}
