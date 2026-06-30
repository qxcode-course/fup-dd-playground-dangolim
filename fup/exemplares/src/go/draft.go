package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	unicos := []int{}

	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		existe := false
		for _, v := range unicos {
			if v == num {
				existe = true
				break
			}
		}

		if !existe {
			unicos = append(unicos, num)
		}
	}

	// Bubble Sort
	for i := 0; i < len(unicos); i++ {
		for j := 0; j < len(unicos)-1-i; j++ {
			if unicos[j] > unicos[j+1] {
				unicos[j], unicos[j+1] = unicos[j+1], unicos[j]
			}
		}
	}

	for i, v := range unicos {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}