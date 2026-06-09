package main

import "fmt"

func main() {
	var a, b, c, d string
	acertos := 0

	fmt.Scan(&a, &b, &c, &d)

	if a == "d" {
		acertos++
	}
	if b == "a" {
		acertos++
	}
	if c == "c" {
		acertos++
	}
	if d == "d" {
		acertos++
	}

	if acertos == 0 {
		fmt.Println("Nunca assistiu")
	} else if acertos == 1 {
		fmt.Println("Ja ouviu falar")
	} else if acertos == 2 {
		fmt.Println("Interessado no assunto")
	} else if acertos == 3 {
		fmt.Println("Fa")
	} else {
		fmt.Println("Super Fa")
	}
}