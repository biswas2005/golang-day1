package operators

import "fmt"

func Assignment() {

	a := 30
	b := 43

	a = b
	fmt.Println(a)
	a += b
	fmt.Println(a)
	a -= b
	fmt.Println(a)
	a *= b
	fmt.Println(a)
	a /= b
	fmt.Println(a)
	a %= b
	fmt.Println(a)
	a <<= b
	fmt.Println(a)
	a >>= b
	fmt.Println(a)

}
