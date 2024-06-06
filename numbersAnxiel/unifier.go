package numbersAnxiel

func Unifier(n int) (str string) {
	if n > 0 && n < 4 {
		str = BelowFour(n)
	} else if n == 4 || n == 5 {
		str = FiveAndFour(n)
	} else if n > 5 && n < 9 {
		str = LessNine(n)
	} else if n == 9 {
		str = Nine(n)
	} else if n == 10 {
		str += Ten(n)
	} else if n > 10 && n < 40 {
		str = UnderThirtyNine(n)
	} else if n >= 40 && n < 50 {
		str += "XL"
		t := n - 40
		str += Unifier(t)
	} else if n >= 50 && n < 90 {
		str += "L"
		b := n - 50
		str += Unifier(b)
	} else if n >= 90 && n < 100 {
		str += "XC"
		b := n - 90
		str += Unifier(b)
	} else if n >= 100 && n < 400 {
		hundreds := n / 100
		str += ThreeHundred(hundreds)
		b := n - (100 * hundreds)
		str += Unifier(b)
	} else if n >= 400 && n < 500 {
		str += "CD"
		b := n - 400
		str += Unifier(b)
	} else if n >= 500 && n < 900 {
		str += "D"
		b := n - 500
		str += Unifier(b)
	} else if n >= 900 && n < 1000 {
		str += "CM"
		b := n - 900
		str += Unifier(b)
	} else if n >= 1000 && n < 4000 {
		thousand := n / 1000
		str += FiveThousand(thousand)
		b := n - (1000 * thousand)
		str += Unifier(b)
	} else if n >= 4000 && n < 5000 {
		str += "MC"
		b := n - 4000
		str += Unifier(b)
	} else if n >= 5000 && n < 9000 {
		str += "_\nV"
		b := n - 5000
		str += Unifier(b)
	} else if n >= 9000 && n < 10000 {
		str += " _\nMX"
		b := n - 9000
		str += Unifier(b)
	} else if n >= 10000 && n < 15000 {
		str += "__\nXV"
		b := n - 10000
		str += Unifier(b)
	} else if n >= 15000 && n < 20000 {
		str += "__\nXV"
		b := n - 15000
		str += Unifier(b)
	} else if n >= 20000 && n < 25000 {
		str += "__\nXV"
		b := n - 20000
		str += Unifier(b)
	} else if n >= 25000 && n < 30000 {
		str += "__\nXV"
		b := n - 25000
		str += Unifier(b)
	} else if n >= 30000 && n < 35000 {
		str += "__\nXV"
		b := n - 30000
		str += Unifier(b)
	} else if n >= 35000 && n < 40000 {
		str += "__\nXV"
		b := n - 35000
		str += Unifier(b)
	} else if n >= 40000 && n < 45000 {
		str += "__\nXL"
		b := n - 40000
		str += Unifier(b)
	} else if n >= 45000 && n < 49000 {
		str += "___\nXLV"
		b := n - 45000
		str += Unifier(b)
	} else if n >= 49000 && n < 50000 {
		str += "__ _\nXLMX"
		b := n - 49000
		str += Unifier(b)
	} else if n >= 50000 && n < 55000 {
		str += "_\nL"
		b := n - 50000
		str += Unifier(b)
	} else if n >= 55000 && n < 60000 {
		str += "__\nLV"
		b := n - 55000
		str += Unifier(b)
	} else if n >= 60000 && n < 65000 {
		str += "__\nLX"
		b := n - 60000
		str += Unifier(b)
	} else if n >= 65000 && n < 70000 {
		str += "___\nLXV"
		b := n - 65000
		str += Unifier(b)
	} else if n >= 70000 && n < 75000 {
		str += "__\nLXX"
		b := n - 70000
		str += Unifier(b)
	} else if n >= 75000 && n < 80000 {
		str += "____\nLXXV"
		b := n - 75000
		str += Unifier(b)
	} else if n >= 80000 && n < 85000 {
		str += "____\nLXXX"
		b := n - 80000
		str += Unifier(b)
	} else if n >= 85000 && n < 90000 {
		str += "_____\nLXXXV"
		b := n - 85000
		str += Unifier(b)
	} else if n >= 90000 && n < 95000 {
		str += "__\nXC"
		b := n - 90000
		str += Unifier(b)
	} else if n >= 95000 && n < 100000 {
		str += "___\nXCV"
		b := n - 95000
		str += Unifier(b)
	} else if n >= 100000 && n < 105000 {
		str = "_\nC"
		b := n - 100000
		str += Unifier(b)
	} else if n == 105000 {
		str = "__\nCV"
	} else if n > 105000 {
		str = `
	for the numbers above 105,000, we can provide you with the following, considering that 1,000,000
	is the highest number possible to be represented by roman numerals as a special character;
				                _
				500,000 ---->   D
				                _
				1,000,000 ----> M
`
	}
	return
}
