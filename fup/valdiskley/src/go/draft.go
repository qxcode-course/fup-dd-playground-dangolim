package main

import "fmt"

func main() {
	var letra rune
	var rotacao int

	fmt.Scanf("%c", &letra)
	fmt.Scan(&rotacao)

	pos := int(letra - 'a')

	pos = (pos + rotacao) % 26

	if pos < 0 {
		pos += 26
	}

	fmt.Printf("%c\n", rune(pos)+'a')
}