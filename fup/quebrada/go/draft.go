package main

import (
	"fmt"
)

func main() {
	var n1, n2 int
	_, err := fmt.Scan(&n1, &n2)
	if err != nil {
		return
	}

	divInteira := n1 / n2
	resto := n1 % n2
	divQuebrada := float64(n1) / float64(n2)

	fmt.Println(divInteira)
	fmt.Println(resto)
	fmt.Printf("%.2f\n", divQuebrada)
}