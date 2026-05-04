package main

import (
	"fmt"
)

func main() {
	var salario float64
	fmt.Scan(&salario)

	var aumento float64

	if salario <= 1000 {
		aumento = 0.20
	} else if salario <= 1500 {
		aumento = 0.15
	} else if salario <= 2000 {
		aumento = 0.10
	} else {
		aumento = 0.05
	}

	novoSalario := salario * (1 + aumento)

	fmt.Printf("%.2f\n", novoSalario)
}