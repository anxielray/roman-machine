package numbersAnxiel

func OverFlow(n int) (str string) {
	if n <= 0 {
		str = `
               there don't exist Roman Numerals for numbers equal to or less than 0
               Similarly, letters don't count as numbers hence cannot be converted to Roman Numerals...
               try numbers such as 400, 600, 1, 0r even +144
`
	}
	return
}
