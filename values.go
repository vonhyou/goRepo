package main

import "fmt"

func main() {
	fmt.Println(2)

	// Wrong! Vars in go will not convert implicitly
	// var a int64 = 3.2

	var a = 10
	fmt.Println(a)

	b, c := 1.2, 3
	// invalid operation: b * c (mismatched types float64 and int)
	// fmt.Println(b * c)
	fmt.Println(b, "x", float64(c), "=", b*float64(c))

	// division by zero
	// fmt.Println(1 / 0)
}
