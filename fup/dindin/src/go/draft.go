package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	c, l := 0, 0
	m, t := 0, 0

	var sabor, turno string

	for i := 0; i < n; i++ {
		fmt.Scan(&sabor, &turno)

		if sabor == "c" {
			c++
		} else {
			l++
		}

		if turno == "m" {
			m++
		} else {
			t++
		}
	}

	if c > l {
		fmt.Println("c")
	} else if l > c {
		fmt.Println("l")
	} else {
		fmt.Println("empate")
	}

	if m < t {
		fmt.Println("m")
	} else if t < m {
		fmt.Println("t")
	} else {
		fmt.Println("empate")
	}
}