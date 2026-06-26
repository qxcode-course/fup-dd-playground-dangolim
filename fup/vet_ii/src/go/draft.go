package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	v := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	fmt.Print("[ ")
	for i := 0; i < n; i++ {
		fmt.Print(v[i], " ")
	}
	fmt.Println("]")
}