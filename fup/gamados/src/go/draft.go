package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)

	p := strings.Split(s, " ")

	for i := 0; i < len(p)-1; i++ {
		if p[i] > p[i+1] {
			fmt.Println("nao")
			return
		}
	}

	fmt.Println("sim")
}