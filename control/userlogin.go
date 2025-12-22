package control

import "fmt"

var a = "Biswas"
var b = "2233"

func User() {
	var username string
	var password string
	fmt.Println("Enter the username: ")
	fmt.Scan(&username)
	if username == a {
		fmt.Printf("Enter the password: ")
		fmt.Scan(&password)
		if password == b {
			fmt.Println("Login successful")
		} else {
			fmt.Println("Wrong password")
		}

	} else {
		fmt.Println("Invalid username")
	}
}
