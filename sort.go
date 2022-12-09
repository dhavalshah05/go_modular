package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Channel1 struct {
	Number   string
	Name     string
	Language string
	Genre    string
	App      string
}

func main() {
	fileBytes, err := os.ReadFile("channels.json")
	if err != nil {
		panic(err)
	}

	var channels []Channel1
	err = json.Unmarshal(fileBytes, &channels)
	if err != nil {
		panic(err)
	}

	sortChannels(channels, false)

	for _, channel := range channels {
		fmt.Println(channel)
	}

}

func sortChannels(channels []Channel1, desc bool) {
	sort.Slice(channels, func(i, j int) bool {
		num1, _ := strconv.Atoi(channels[i].Number)
		num2, _ := strconv.Atoi(channels[j].Number)
		if desc {
			return num1 > num2
		} else {
			return num1 < num2
		}
	})
}
