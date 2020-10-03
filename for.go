package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		// should replace i += 1 with i++
		// i += 1
		// No ++i
		// ++i
		i++
	}

	// syntax error: unexpected newline, expecting { after for clause
	// {} cannot be omitted
	for j := 1; j < 10; j++ {
		fmt.Println(j, ":", j*j)
	}
	// syntax error: var declaration not allowed in for initializer
	// for var k int = 1; k < 5; k++

	for {
		fmt.Println("嗷嗷！！")
		break
	}
}
