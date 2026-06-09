package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	texto, _ := reader.ReadString('\n')
	texto = strings.TrimSpace(texto)

	palavras := strings.Fields(texto)

	fmt.Println(strings.Join(palavras, " "))
}