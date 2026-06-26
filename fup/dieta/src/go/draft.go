package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)

	soma := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		soma += x
	}

	media := float64(soma) / float64(n)
	fmt.Printf("%.1f\n", media)
}