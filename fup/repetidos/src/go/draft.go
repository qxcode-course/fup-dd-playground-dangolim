package main

import "fmt"

func main() {
	var p, n, x int
	fmt.Scan(&p, &n)

	cont := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x == p {
			cont++
		}
	}

	fmt.Println(cont)
}