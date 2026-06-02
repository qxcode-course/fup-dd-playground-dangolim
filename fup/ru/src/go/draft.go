package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	s, _ := in.ReadString('\n')

	v, c := "", ""

	for _, l := range s {
		if l == 'a' || l == 'e' || l == 'i' || l == 'o' || l == 'u' {
			v += string(l)
		} else if l >= 'a' && l <= 'z' {
			c += string(l)
		}
	}

	fmt.Println(v)
	fmt.Println(c)
}