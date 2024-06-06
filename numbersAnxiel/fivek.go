package numbersAnxiel

func FiveThousand(n int)(s string) {
	for x := 1; x <= n; x++ {
		s += "M"
	}
	return
}