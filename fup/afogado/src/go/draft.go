package main

import "fmt"

func consegue(P, E, S int) bool {
	pos, jump := 0, S

	for jump >= 0 {
		if pos+jump >= P {
			return true
		}

		pos += jump
		pos -= E

		if pos < 0 {
			return false
		}

		jump -= 10
	}
	return false
}

func main() {
	var P, E int
	fmt.Scan(&P, &E)

	for s := 0; ; s++ {
		if consegue(P, E, s) {
			fmt.Println(s)
			break
		}
	}
}