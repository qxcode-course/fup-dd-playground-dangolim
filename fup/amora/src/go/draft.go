package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	texto, _ := in.ReadString('\n')
	trecho, _ := in.ReadString('\n')

	texto = strings.TrimSpace(texto)
	trecho = strings.TrimSpace(trecho)

	fmt.Println(strings.Count(texto, trecho))
}