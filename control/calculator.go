package control

import (
	"fmt"
	"math"
)

func Calci() {
	var a, b float64
	var o string

	fmt.Print("")
	fmt.Scan(&a, &o, &b)

	if o == "+" {
		fmt.Print(a + b)
	} else if o == "-" {
		fmt.Print(a - b)
	} else if o == "*" {
		fmt.Print(a * b)
	} else if o == "/" {
		fmt.Print(a / b)
	} else if o == "%" {
		fmt.Print(math.Mod(a, b))
	}

}
