package main

import "fmt"

func main() {
	var valor float64
	fmt.Scan(&valor)
	total := int(valor*100 + 0.5)
	opcoes := []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 25, 10, 5}

	for _, moeda := range opcoes {
		qtd := total / moeda
		if qtd > 0 {
			fmt.Printf("%d de %.2f\n", qtd, float64(moeda)/100)
		}
		total %= moeda
	}

	if total > 0 {
		fmt.Printf("Falta %.2f\n", float64(total)/100)
	}
}