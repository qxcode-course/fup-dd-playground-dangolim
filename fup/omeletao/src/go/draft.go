package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	maior := a

	if b > maior {
		maior = b
	}
	if c > maior {
		maior = c
	}
	if d > maior {
		maior = d
	}

	fmt.Println(maior)
}