package numbersAnxiel

func UnderThirtyNine(num int) (str string) {
	tens := num / 10
	for x := 1; x <= tens; x++ {
		str += "X"
	}
	ones := num % 10
	if ones < 4 && ones > 0 {
		str += BelowFour(ones)
	} else if ones == 4 || ones == 5 {
		str += FiveAndFour(ones)
	} else if ones > 5 && ones < 9 {
		str += LessNine(ones)
	} else if ones == 9 {
		str += Nine(ones)
	} else {
		str += ""
	}
	return
}
