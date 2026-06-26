package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)

	v1 := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&v1[i])
	}

	fmt.Scan(&m)

	v2 := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&v2[i])
	}

	for i := 0; i < n; i++ {
		achou := false

		for j := 0; j < m; j++ {
			if v1[i] == v2[j] {
				achou = true
				break
			}
		}

		if !achou {
			fmt.Println("nao")
			return
		}
	}

	fmt.Println("sim")
}