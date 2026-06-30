package main

import "fmt"

func main() {
	var n, inferior, superior int
	fmt.Scan(&n, &inferior, &superior)

	contador := 0

	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		if num >= inferior && num <= superior {
			contador++
		}
	}

	fmt.Println(contador)
}