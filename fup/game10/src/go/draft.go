package main

import (
	"fmt"
)

func main() {
	var N, D, A int

	fmt.Scan(&N)
	fmt.Scan(&D)
	fmt.Scan(&A)

	if D >= A {
		fmt.Println(D - A)
	} else {
		fmt.Println((N - A) + D)
	}
}
