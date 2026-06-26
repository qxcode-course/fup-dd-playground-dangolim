package main

import "fmt"

func main() {
	var P, S, E int
	fmt.Scan(&P, &S, &E)

	pos, salto := 0, S

	for {
		top := pos + salto

		if top >= P {
			fmt.Printf("%d saiu\n", pos)
			break
		}

		fmt.Printf("%d %d\n", pos, top)

		salto -= 10
		if salto < 0 { // não consegue mais saltar
			break
		}

		pos = top - E
		if pos < 0 {
			fmt.Printf("%d morreu\n", pos)
			break
		}
	}
}