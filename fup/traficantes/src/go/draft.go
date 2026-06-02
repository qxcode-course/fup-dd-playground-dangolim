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
	antiga, _ := in.ReadString('\n')
	nova, _ := in.ReadString('\n')

	texto = strings.TrimSpace(texto)
	antiga = strings.TrimSpace(antiga)
	nova = strings.TrimSpace(nova)

	res := ""

	for i := 0; i < len(texto); {
		if i+len(antiga) <= len(texto) && texto[i:i+len(antiga)] == antiga {
			res += nova
			i += len(antiga)
		} else {
			res += string(texto[i])
			i++
		}
	}

	fmt.Println(res)
}