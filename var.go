package main

import "fmt"

func main() {
	var s string = "hello, world!"
	fmt.Println(s)
	var a, b int = 1, 2
	fmt.Println(a + b)
	c := 3.14
	fmt.Println(c)

	// Wrong way!
	// var x float32, y int = 1.0, 233
	x, y := 1.1, 233
	fmt.Println(x, y)
}
