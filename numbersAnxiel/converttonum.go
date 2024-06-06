package numbersAnxiel

import (
	"fmt"
	"os"
	"strconv"
)

func ArgsToNum(s string) (n int) {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("your string contains a character that could not be converted to a number...")
	}
	return
}
