package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	animais := make([]int, n)
	usado := make([]bool, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&animais[i])
	}

	casais := 0

	for i := 0; i < n; i++ {
		if usado[i] {
			continue
		}

		for j := i + 1; j < n; j++ {
			if !usado[j] && animais[i] == -animais[j] {
				casais++
				usado[i] = true
				usado[j] = true
				break
			}
		}
	}

	fmt.Println(casais)
}