package numbersAnxiel

func BelowFour(n int) (str string) {
	for x := 1; x <= n; x++ {
		str += "I"
	}
	return
}

func FiveAndFour(n int) string {
	var str string
	if n == 4 {
		str += "IV"
	} else if n == 5 {
		str += "V"
	}
	return str
}

func LessNine(n int) (str string) {
	str += "V"
	str += BelowFour(n - 5)
	return
}

func Nine(n int) (str string) {
	str += "IX"
	return
}

func Ten(n int) (str string) {
	str += "X"
	return
}
