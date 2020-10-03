package main

import (
	"fmt"
	"math"
)

const str string = "It's a constant"

func main() {
	fmt.Println(str)

	const n = 500000
	const f = 3e20 / n

	fmt.Println(f)
	fmt.Println(int64(f))

	fmt.Println(math.Pi, math.Sin(math.Pi/2))
}
