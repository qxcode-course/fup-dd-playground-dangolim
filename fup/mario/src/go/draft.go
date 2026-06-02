package main

import "fmt"

func main() {
	var n, maior int
	fmt.Scan(&n)

	v := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])

		if v[i] > maior {
			maior = v[i]
		}
	}

	for i := maior; i >= 1; i-- {
		for j := 0; j < n; j++ {
			if v[j] >= i {
				fmt.Print("#")
			} else {
				fmt.Print("_")
			}
		}
		fmt.Println()
	}
}