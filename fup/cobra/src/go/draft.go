package main

import (
	"fmt"
)

func main() {
	var N, X, Y, S int
	var C string

	fmt.Scan(&N)
	fmt.Scan(&X)
	fmt.Scan(&Y)
	fmt.Scan(&C)
	fmt.Scan(&S)

	switch C {
	case "U":
		Y = (Y - S%N + N) % N
	case "D":
		Y = (Y + S) % N
	case "L":
		X = (X - S%N + N) % N
	case "R":
		X = (X + S) % N
	}

	fmt.Println(X, Y)
}
