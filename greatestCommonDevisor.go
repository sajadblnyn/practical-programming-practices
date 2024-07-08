package main

import "fmt"

func main() {
	x := 0
	y := 0
	fmt.Scan(&x, &y)
	r := recursive_bmm(x, y)

	fmt.Println(r)
}

func recursive_bmm(x int, y int) int {
	if y <= 0 {
		return x
	}

	return recursive_bmm(y, (x % y))
}
