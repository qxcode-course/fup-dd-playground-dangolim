package main

import "fmt"

func main() {
	var L, R, D int

	fmt.Scan(&L)
	fmt.Scan(&R)
	fmt.Scan(&D)

	if (R > 50) && (L < R) && (R > D) {
		fmt.Println("S")
	} else {
		fmt.Println("N")
	}
}