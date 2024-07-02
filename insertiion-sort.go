package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scanln(&n)

	var in int

	var nums []int
	for i := 0; i < n; i++ {
		fmt.Scan(&in)

		nums = append(nums, in)
	}

	var p int
	var c int

	for i := 1; i < n; i++ {
		p = i
		for true {
			if (p <= 0) || nums[p] >= nums[p-1] {
				break
			}
			c = nums[p]
			nums[p] = nums[p-1]
			nums[p-1] = c
			p--
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", nums[i])
	}

}
