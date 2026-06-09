package main

import "fmt"

func main() {
	var A, G, Ra, Rg float64

	fmt.Scan(&A, &G, &Ra, &Rg)

	if A/Ra < G/Rg {
		fmt.Println("A")
	} else {
		fmt.Println("G")
	}
}