package main

import "fmt"

func main() {
	var L, C int

	fmt.Scan(&L, &C)

	if (L+C)%2 == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}