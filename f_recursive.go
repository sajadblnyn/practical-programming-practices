package main

import "fmt"

func main() {
	x := 0

	fmt.Scanln(&x)

	r := recursive(x)

	fmt.Println(r)
}

func recursive(x int) int {
	if x == 0 {
		return 5
	}

	tmp := recursive(x - 1)
	if (x % 2) == 0 {
		return tmp - 21
	} else {
		return tmp * tmp
	}
}
