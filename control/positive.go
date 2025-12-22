package control

import "fmt"

func If() {
	var a int
	fmt.Print("Enter a number: ")
	fmt.Scan(&a)
	if a > 0 {
		fmt.Println("It is a Positive number.")
	} else {
		fmt.Println("It is a negative number.")
	}
}
