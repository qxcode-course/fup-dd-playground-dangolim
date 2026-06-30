package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	somaJedi := 0
	somaSith := 0

	for i := 0; i < t; i++ {
		var forca int
		fmt.Scan(&forca)

		if i < t/2 {
			somaJedi += forca
		} else {
			somaSith += forca
		}
	}

	if somaJedi > somaSith {
		fmt.Println("Jedi")
	} else if somaSith > somaJedi {
		fmt.Println("Sith")
	} else {
		fmt.Println("Empate")
	}
}