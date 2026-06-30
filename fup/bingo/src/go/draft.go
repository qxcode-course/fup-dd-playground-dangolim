package main

import "fmt"

func main() {
	// Cartela fixa da avó
	cartela := map[int]bool{
		1: true, 9: true, 27: true, 23: true,
		34: true, 20: true, 37: true, 47: true,
		30: true, 87: true, 55: true, 69: true,
		13: true, 60: true, 99: true, 66: true,
	}

	contador := 0

	for i := 0; i < 6; i++ {
		var num int
		fmt.Scan(&num)

		if cartela[num] {
			contador++
		}
	}

	fmt.Println(contador)
}