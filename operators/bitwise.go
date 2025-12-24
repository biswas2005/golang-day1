package operators

import "fmt"

func Assign() {
	a := 10
	b := 20

	result1 := a & b
	fmt.Printf("a&b: %d\n", result1)
	result2 := a | b
	fmt.Printf("a|b: %d\n", result2)
	result3 := a ^ b
	fmt.Printf("a^b: %d\n", result3)
	result4 := a << b
	fmt.Printf("a<<b: %d\n", result4)
	result5 := a >> b
	fmt.Printf("a>>b: %d\n", result5)
}
