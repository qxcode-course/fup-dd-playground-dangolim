package main

import "fmt"

func main() {
	var p, s, e, x int
	fmt.Scan(&p, &s, &e)

	for x+s < p {
		fmt.Println(x, x+s)
		x += s - e
	}

	fmt.Println(x, "saiu")
}