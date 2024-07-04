package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int

	fmt.Scanln(&n)

	var s int

	c := len(strconv.Itoa(n))
	for i := 0; i < c; i++ {
		s = s + (n % 10)
		n = n / 10

	}
	fmt.Println(s)
}
