package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 5, 7}))
}

func maxProfit(prices []int) (maxProf int) {
	buyPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < buyPrice {
			buyPrice = prices[i]
		}

		if (prices[i] - buyPrice) > maxProf {
			maxProf = prices[i] - buyPrice
		}
	}
	return maxProf
}
