package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	name := getStringFromUser("What is your name?")
	age, err := getIntFromUser("What is your age?")
	if err != nil {
		fmt.Println("Please enter valid age!!!")
		return
	}

	fmt.Println(name, age)
}

func getIntFromUser(title string) (vaule int, err error) {
	stringVal := getStringFromUser(title)
	result, err := strconv.Atoi(stringVal)
	return result, err
}

func getStringFromUser(title string) string {
	fmt.Println(title, ":", " ")
	reader := bufio.NewReader(os.Stdin)
	result, _ := reader.ReadString('\n')
	return strings.TrimSpace(result)
}
