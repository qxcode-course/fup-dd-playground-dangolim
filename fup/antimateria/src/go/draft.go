package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b string

	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)

	maior := 0

	// Busca a maior sobreposição entre
	// o final de a e o começo de b
	limite := len(a)
	if len(b) < limite {
		limite = len(b)
	}

	for i := 1; i <= limite; i++ {

		if a[len(a)-i:] == b[:i] {
			maior = i
		}
	}

	// Remove a parte aniquilada dos dois lados
	resultado := a[:len(a)-maior] + b[maior:]

	fmt.Println(resultado)
}