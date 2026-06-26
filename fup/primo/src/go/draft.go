package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n < 2 {
		fmt.Println(0)
		return
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Println(0)
			return
		}
	}

	fmt.Println(1)
}