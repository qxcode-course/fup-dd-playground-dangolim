package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}

	fmt.Print("[ ")
	for i := n - 1; i >= 0; i-- {
		fmt.Print(vetor[i], " ")
	}
	fmt.Println("]")
}