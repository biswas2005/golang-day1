package day1

import "fmt"

const (
	a = iota
	b
)
const (
	c = iota - 3
	d
)

func Cons() {
	fmt.Println(a, b, c, d)
}
