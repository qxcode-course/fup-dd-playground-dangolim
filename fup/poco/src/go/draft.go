package main

import "fmt"

func main() {
	var P, S, E int
	fmt.Scan(&P, &S, &E)

	pos := 0

	for {
		fmt.Printf("%d %d\n", pos, pos+S)

		pos += S
		if pos >= P {
			break
		}

		pos -= E
		if pos < 0 {
			break
		}

		S -= 10
	}
}