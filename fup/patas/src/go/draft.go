package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var chico, cebolinha, n int
	fmt.Scan(&chico, &cebolinha, &n)

	total := 0
	var animal string

	for i := 0; i < n; i++ {
		fmt.Scan(&animal)

		if animal == "g" {
			total += 2
		} else {
			total += 4
		}
	}

	fmt.Println(total)

	d1 := abs(chico - total)
	d2 := abs(cebolinha - total)

	if d1 < d2 {
		fmt.Println("Chico Bento")
	} else if d2 < d1 {
		fmt.Println("Cebolinha")
	} else {
		fmt.Println("empate")
	}
}