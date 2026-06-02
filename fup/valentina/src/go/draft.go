package main

import (
	"fmt"
)

func main() {
	var c1, op, c2 rune

	fmt.Scanf("%c\n", &c1)
	fmt.Scanf("%c\n", &op)
	fmt.Scanf("%c", &c2)

	v1 := int(c1 - 'a')
	v2 := int(c2 - 'a')

	var resultado int

	if op == '+' {
		resultado = (v1 + v2) % 26
	} else if op == '-' {
		resultado = (v1 - v2 + 26) % 26
	}

	fmt.Printf("%c\n", rune(resultado)+'a')
}