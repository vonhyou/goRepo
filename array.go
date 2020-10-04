package main

import "fmt"

func main() {
	var a [5]int
	a[4] = 1
	fmt.Println("emp:", a, a[4])
	// type [5]int is not an expression
	// b := [5]int{0}
	b := [5]int{0, 1, 2, 3, 4}
	fmt.Println("arr:", b)

	var tda [4][3]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			tda[i][j] = i + j
		}
	}
	fmt.Println("2D Array:", tda)
}
