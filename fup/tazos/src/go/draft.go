package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	v := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	resp := []int{}
	cont, maxCont := 1, 0

	for i := 1; i <= n; i++ {
		if i < n && v[i] == v[i-1] {
			cont++
		} else {
			if cont > maxCont {
				maxCont = cont
				resp = []int{v[i-1]}
			} else if cont == maxCont {
				resp = append(resp, v[i-1])
			}
			cont = 1
		}
	}

	fmt.Print("[ ")
	for _, x := range resp {
		fmt.Print(x, " ")
	}
	fmt.Println("]")
}