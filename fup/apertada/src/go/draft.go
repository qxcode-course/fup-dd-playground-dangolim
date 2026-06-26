package main

import "fmt"

func main() {
	var x, menor int

	fmt.Scan(&menor)

	for i := 0; i < 4; i++ {
		fmt.Scan(&x)
		if x < menor {
			menor = x
		}
	}

	fmt.Println(menor)
}