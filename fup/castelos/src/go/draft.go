package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i*i <= n; i++ {
		if i*i == n {
			fmt.Println("sim")
			return
		}
	}

	fmt.Println("nao")
}