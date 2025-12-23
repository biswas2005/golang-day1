package control

import (
	"fmt"
	"math"
)

func Calci() {
	var value1, value2 float64
	var operator string

	fmt.Println("enter value1 <operator> value2")
	fmt.Scan(&value1, &operator, &value2)

	if operator == "+" {
		fmt.Print(value1 + value2)
	} else if operator == "-" {
		fmt.Print(value1 - value2)
	} else if operator == "*" {
		fmt.Print(value1 * value2)
	} else if operator == "/" {
		fmt.Print(value1 / value2)
	} else if operator == "%" {
		fmt.Print(math.Mod(value1, value2))
	}

}
