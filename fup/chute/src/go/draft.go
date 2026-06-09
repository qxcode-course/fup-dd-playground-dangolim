package main

import "fmt"

func main() {
	var valor, primeiro, segundo int

	fmt.Scan(&valor)
	fmt.Scan(&primeiro)
	fmt.Scan(&segundo)

	dist1 := valor - primeiro
	if dist1 < 0 {
		dist1 = -dist1
	}

	dist2 := valor - segundo
	if dist2 < 0 {
		dist2 = -dist2
	}

	if dist1 < dist2 {
		fmt.Println("primeiro")
	} else if dist2 < dist1 {
		fmt.Println("segundo")
	} else {
		fmt.Println("empate")
	}
}