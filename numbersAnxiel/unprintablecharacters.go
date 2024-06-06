package numbersAnxiel

import (
	"os"
	"strings"
)

func Unprintable(s string) (str string) {
	for _, char := range s {
		if !((char >= 48 && char <= 57) || (rune(s[0]) == '-' || rune(s[0]) == '+')) {
			str = "your string contains a non-number character!..."
			os.Exit(0)
		} else if rune(s[0]) == '+' {
			str = strings.TrimLeft(s, "+")
		} else {
			str = s
		}
	}
	return
}
