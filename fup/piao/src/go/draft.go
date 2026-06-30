package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var L, N int
	fmt.Scan(&L, &N)

	ganhador := -1
	melhorDist := L + 1 // maior que qualquer distância válida

	perdedor := 0
	piorDist := -1

	for i := 0; i < N; i++ {
		var jogada int
		fmt.Scan(&jogada)

		dist := abs(jogada)

		// Verifica ganhador
		if dist <= L {
			if dist <= melhorDist { // <= garante último em empate
				melhorDist = dist
				ganhador = i
			}
		}

		// Verifica perdedor
		if dist >= piorDist { // >= garante último em empate
			piorDist = dist
			perdedor = i
		}
	}

	if ganhador == -1 {
		fmt.Println("nenhum")
	} else {
		fmt.Println(ganhador)
	}

	fmt.Println(perdedor)
}