package operators

import "fmt"

func Logic() {
	a := 20
	b := 10

	if a == b && a <= b {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

}
