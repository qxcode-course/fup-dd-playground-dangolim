package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)

	soldados, rebeldes := 0, 0

	for i := 0; i < n; i++ {
		fmt.Scan(&x)

		if x%2 == 0 {
			rebeldes += x
		} else {
			soldados += x
		}
	}

	if soldados > rebeldes {
		fmt.Println("soldados")
	} else if rebeldes > soldados {
		fmt.Println("rebeldes")
	} else {
		fmt.Println("empate")
	}
}