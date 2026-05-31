package main

import "fmt"

func inverter(n int) int {
	inv := 0

	for n > 0 {
		inv = inv*10 + n%10
		n /= 10
	}

	return inv
}

func main() {
	var n int
	fmt.Scan(&n)

	if n == inverter(n) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}