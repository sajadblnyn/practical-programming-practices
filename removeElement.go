package main

import "fmt"

func main() {
	fmt.Println()
}

func removeElement(nums []int, val int) int {
	var k int

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			k++
		}
	}
	return k
}
