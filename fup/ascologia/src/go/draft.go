package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	nome, _ := reader.ReadString('\n')
	nome = strings.TrimSpace(nome)

	soma := 0

	for _, c := range nome {
		soma += int(c)
	}

	fmt.Println(soma % 50)
}