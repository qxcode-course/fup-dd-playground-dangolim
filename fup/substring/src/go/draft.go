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

	var i, q int
	fmt.Fscan(in, &i, &q)

	if i < 0 || i >= len(s) || q <= 0 {
		fmt.Println("")
		return
	}

	f := i + q
	if f > len(s) {
		f = len(s)
	}

	fmt.Println(s[i:f])
}
