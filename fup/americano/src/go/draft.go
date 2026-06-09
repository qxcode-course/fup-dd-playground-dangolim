package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	soma := a + b + c + d

	if soma == 0 {
		fmt.Println("nenhum")
	} else {
		r := soma % 4

		if r == 1 {
			fmt.Println("jog1")
		} else if r == 2 {
			fmt.Println("jog2")
		} else if r == 3 {
			fmt.Println("jog3")
		} else {
			fmt.Println("jog4")
		}
	}
}