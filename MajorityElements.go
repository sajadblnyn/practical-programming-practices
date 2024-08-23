package main

import "fmt"

// when we use this algorithm  to find the number that count of that number is more than Half of the count of all numbers
func main() {
	fmt.Println(majorityElement([]int{3, 5, 3, 2, 1}))
}

func majorityElement(nums []int) int {
	var r int
	var c int

	for _, v := range nums {
		if r == 0 {
			c = v
		}

		if v == c {
			r++
		} else {
			r--
		}
	}
	return c
}
