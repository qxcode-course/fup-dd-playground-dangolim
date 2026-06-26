package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println(0)
		return
	}
	if n <= 2 {
		fmt.Println(1)
		return
	}

	var a, b int64 = 1, 1

	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}

	fmt.Println(b)
}