package main

import "fmt"

func main() {
	var n, num int
	var impares []int
	var pares []int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&num)

		if num%2 == 0 {
			pares = append(pares, num)
		} else {
			impares = append(impares, num)
		}
	}

	fmt.Print("[ ")
	for _, v := range impares {
		fmt.Print(v, " ")
	}
	fmt.Println("]")

	fmt.Print("[ ")
	for _, v := range pares {
		fmt.Print(v, " ")
	}
	fmt.Println("]")
}