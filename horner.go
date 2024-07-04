package main

import (
	"fmt"
)

const mod = 1000000007

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	coeffs := make([]int, n+1)
	for i := 0; i <= n; i++ {
		fmt.Scan(&coeffs[i])
	}

	res := 0
	for i := 0; i <= n; i++ {
		res = (res*x + coeffs[i]) % mod
	}

	if res < 0 {
		res += mod
	}

	fmt.Println(res)
}
