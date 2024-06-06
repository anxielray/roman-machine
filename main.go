package main

import (
	"fmt"
	roman "roman/numbersAnxiel"
)

func main() {
	fmt.Print("feed me a number: ")
	var num int
	fmt.Scanf("%v", &num)
	fmt.Println()
	if num <= 0 || num > 105000 {
		fmt.Println(roman.OverFlow(num))
	} else {
		fmt.Printf("%v in Roman Numerals is\n%v", num, roman.Unifier(num))
	}
}
