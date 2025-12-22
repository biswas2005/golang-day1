package control

import "fmt"

func Evenodd() {
	var a int
	fmt.Print("Enter a number: ")
	fmt.Scan(&a)
	if a%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
