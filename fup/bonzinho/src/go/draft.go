package main

import (
	"fmt"
)

func main() {
	var p1, p2, p3, t float64
	_, err := fmt.Scan(&p1, &p2, &p3, &t)
	if err != nil {
		return
	}

	menor := p1
	if p2 < menor {
		menor = p2
	}
	if p3 < menor {
		menor = p3
	}

	media := ((p1 + p2 + p3) - menor + t) / 3.0

	if media >= 7.0 {
		fmt.Printf("Aprovado com %.1f\n", media)
	} else {
		fmt.Printf("Final com %.1f\n", media)
	}
}