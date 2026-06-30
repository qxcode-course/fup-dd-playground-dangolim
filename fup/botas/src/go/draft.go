package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// Como os tamanhos vão de 30 a 60, dá pra usar arrays
	esquerda := make([]int, 61)
	direita := make([]int, 61)

	for i := 0; i < n; i++ {
		var tamanho int
		var lado string

		fmt.Scan(&tamanho, &lado)

		if lado == "E" {
			esquerda[tamanho]++
		} else {
			direita[tamanho]++
		}
	}

	pares := 0

	for tamanho := 30; tamanho <= 60; tamanho++ {
		if esquerda[tamanho] < direita[tamanho] {
			pares += esquerda[tamanho]
		} else {
			pares += direita[tamanho]
		}
	}

	fmt.Println(pares)
}