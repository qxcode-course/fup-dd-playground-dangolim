package main

import "fmt"

func main() {
	var total, qtd int
	fmt.Scan(&total, &qtd)

	figs := make([]int, qtd)
	for i := 0; i < qtd; i++ {
		fmt.Scan(&figs[i])
	}

	repetidas := []int{}
	presente := make([]bool, total+1)

	for i := 0; i < qtd; i++ {
		if i > 0 && figs[i] == figs[i-1] {
			repetidas = append(repetidas, figs[i])
		}
		if figs[i] >= 1 && figs[i] <= total {
			presente[figs[i]] = true
		}
	}

	fmt.Print("[ ")
	for _, x := range repetidas {
		fmt.Print(x, " ")
	}
	fmt.Println("]")

	fmt.Print("[ ")
	for i := 1; i <= total; i++ {
		if !presente[i] {
			fmt.Print(i, " ")
		}
	}
	fmt.Println("]")
}