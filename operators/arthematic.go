package operators

import "fmt"

func Ops() {
	a := 10
	b := 20

	result1 := a + b
	fmt.Printf("The sum is: %d+%d=%d\n", a, b, result1)

	result2 := a - b
	fmt.Printf("The sub is: %d-%d=%d\n", a, b, result2)

	result3 := a * b
	fmt.Printf("The mul is: %d*%d=%d\n", a, b, result3)

	result4 := a / b
	fmt.Printf("The div is: %d/%d=%d\n", a, b, result4)

	result5 := a % b
	fmt.Printf("The mod is: %d %% %d=%d\n", a, b, result5)
}
