package main

import "fmt"

func main() {
	var c1, l1, c2, l2 int

	fmt.Scan(&c1, &l1, &c2, &l2)

	a1 := c1 * l1
	a2 := c2 * l2

	if a1 > a2 {
		fmt.Println(a1)
	} else {
		fmt.Println(a2)
	}
}