package main

import "fmt"

func main() {
	var result = 0
	add(1, 3, &result)
	fmt.Println("Result is ", result)
}

func add(num1 int, num2 int, result *int) {
	*result = num1 + num2
}
