package main

import "fmt"

func main() {

	l1 := []int{1, 2, 4, 5, 6, 0}
	l2 := []int{3}
	m := 5
	n := 1

	fmt.Printf("%v\n", merge(l1, m, l2, n))

}

func merge(nums1 []int, m int, nums2 []int, n int) []int {

	nums1 = append(nums1[:m], nums2[:n]...)

	p := m + n - 1
	m = m - 1
	n = n - 1

	for i := p; i >= 0; i-- {
		if m < 0 {

			// nums1[i] = nums2[n]
			// n--

			copy(nums1[:i+1], nums2[:n+1])
			break

		} else if n < 0 {
			break
		} else if nums1[m] >= nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]

			n--
		}
	}

	return nums1
}
