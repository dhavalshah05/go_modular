package input

import (
	"bufio"
	"os"
	"strings"
)

func GetStringFromUser() string {
	reader := bufio.NewReader(os.Stdin)
	result, _ := reader.ReadString('\n')
	return strings.TrimSpace(result)
}
