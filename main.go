package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	name := "Variable"
	name = "Value is changed!"
	fmt.Println("Hello", "Dhaval", name)

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error while getting directory")
	}

	fmt.Println(dir)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error while fetching files")
	}

	for index, file := range files {
		fileParts := strings.Split(file.Name(), ".")
		fileExt := fileParts[1]
		fmt.Println(index, file.Name(), file.Size(), file.IsDir(), fileExt)
	}
}
