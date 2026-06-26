package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println()
		return
	}

	v := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	for i := 0; i < n; i++ {
		fmt.Println(v[i])
	}
}