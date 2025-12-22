package control

import "fmt"

func Marks() {
	var a int
	fmt.Print("Enter student mark: ")
	fmt.Scan(&a)
	if a >= 90 {
		fmt.Println("A grade")
	} else if a >= 75 {
		fmt.Println("B grade")
	} else if a >= 50 {
		fmt.Println("C grade")
	} else if a >= 25 {
		fmt.Println("D grade")
	} else {
		fmt.Println("Fail")
	}
}
