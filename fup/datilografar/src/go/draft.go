package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var quebrado string
	var numero string

	fmt.Fscan(reader, &quebrado)
	fmt.Fscan(reader, &numero)

	resultado := ""

	for _, c := range numero {
		if string(c) != quebrado {
			resultado += string(c)
		}
	}

	resultado = strings.TrimLeft(resultado, "0")

	if resultado == "" {
		resultado = "0"
	}

	fmt.Println(resultado)
}