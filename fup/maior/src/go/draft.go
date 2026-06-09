package main

import "fmt"

func main() {
	var chute, valor float64
	var op string

	fmt.Scan(&chute, &op, &valor)

	if chute == valor || (op == "m" && valor > chute) || (op == "M" && valor < chute) {
		fmt.Println("primeiro")
	} else {
		fmt.Println("segundo")
	}
}