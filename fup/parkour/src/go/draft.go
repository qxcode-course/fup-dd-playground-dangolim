package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}

	parkour := 0

	for i := 1; i < n; i++ {
		diff := vetor[i] - vetor[i-1]

		if diff < 0 {
			diff = -diff
		}

		if diff > 1 {
			parkour++
		}
	}

	fmt.Println(parkour)
}