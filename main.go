package main

import (
	"fmt"
	"strconv"
)

func main() {
	intToString()
	intToFloat()
	stringToInt()
	stringToFloat64()
	stringToFloat32()
	floatToString()
	floatToInt()
}

func intToString() {
	age := 28
	converted := strconv.Itoa(age)
	fmt.Printf("%d, %s\n", age, converted)
}

func intToFloat() {
	age := 28
	ageStr := float32(age)
	fmt.Printf("%d, %f\n", age, ageStr)
}

func stringToInt() {
	age := "28"
	converted, err := strconv.Atoi(age)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %d\n", age, converted)
}

func stringToFloat64() {
	age := "28"
	converted, err := strconv.ParseFloat(age, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %f\n", age, converted)
}

func stringToFloat32() {
	age := "28"
	converted, err := strconv.ParseFloat(age, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %f\n", age, converted)
}

func floatToString() {
	amount := 23.50
	converted := fmt.Sprintf("%.2f", amount)
	fmt.Printf("%f, %s\n", amount, converted)
}

func floatToInt() {
	amount := 23.50
	converted := int(amount)
	fmt.Printf("%f, %d\n", amount, converted)
}
