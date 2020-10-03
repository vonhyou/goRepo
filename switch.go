package main

import "fmt"

func main() {
	// non-bool 7 % 2 (type int) used as if condition
	// if 7 % 2 {
	if 7%2 == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}

	if n := 9; n < 0 {
		fmt.Println("n is negative")
	} else if n < 10 {
		fmt.Println("n has only one digital")
	} else {
		fmt.Println("n has multiple digitals")
	}

	i := 2
	fmt.Print("wirte", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) { // declared but not used
		case bool:
			fmt.Println("It's a bool")
		case int:
			fmt.Println("It's a int")
		default:
			fmt.Println("I dont know")
		}
	}
	whatAmI(1)
	whatAmI(true)
	whatAmI("hello")
}
