package main

import (
	"fmt"
	"slices"
	"strconv"
)

func filtrar_impares(nums []int) []int {
	impares := make([]int, 0, len(nums))
	for _, elem := range nums {
		if elem%2 == 1 {
			impares = append(impares, elem)
		}
	}
	return impares
}

func index(num []int, valor int) int {
	for i, elem s:= range nums {
		if elem == valor {
			return i
		}
	}
	return -1
}

func contains(nums []int, valor int) bool {
	for _, elem := range nums {
		if elem == valor {
			return true
		}
	}
	return false
}

func count(nums []int, valor int) int {
	contador := 0
	for _, elem := range nums {
		if elem == valor {
			contador += 1
		}
	}
	return contador
}

func separar_figurinhas(montante []int) ([]int, []int) { //tupla
	album := make([]int, 0, len(montante))
	repet := make([]int, 0, len(montante))
	for _, fig := range montante {
		if !contains(album, fig) {
			album = append(album, fig)
		} else {
			repet = append(repet, fig)
		}
	}
	return album, repet
}

func main() {
	var montante []int = make([]int, 0, 1)
	fmt.Println(montante, len(montante), cap(montante))
	montante = append(montante, 7, 3, 2, 1, 9, 1, 2, 3, 4, 5, 4, 3, 2, 1, 2, 5, 7)
	// album: 1, 2, 3, 4, 5, 7
	// trocar: 4, 3, 2, 1, 2, 5
	num, err := strconv.Atoi("32432")
	if err == nil {
		fmt.Println(num)
	} else {
		fmt.Println(err)
	}
	album, repet := separar_figurinhas(montante)

	slices.Sort(repet)
	fmt.Println(album)
	fmt.Println(repet)
