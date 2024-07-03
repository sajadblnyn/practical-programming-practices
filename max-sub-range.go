package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	max_current := nums[0]
	max_global := nums[0]

	for i := 1; i < n; i++ {
		// if max_current+nums[i] > nums[i] {
		// 	max_current = max_current + nums[i]
		// } else {
		// 	max_current = nums[i]
		// }
		// if max_global < max_current {
		// 	max_global = max_current
		// }
		max_current = max((max_current + nums[i]), nums[i])
		max_global = max(max_global, max_current)

	}
	fmt.Println(max_global)
}
