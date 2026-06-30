package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	v := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	risco := 0

	for i := 0; i < n; i++ {
		if v[i] == 0 {
			temMedico := false

			if i > 0 && v[i-1] == 1 {
				temMedico = true
			}

			if i < n-1 && v[i+1] == 1 {
				temMedico = true
			}

			if !temMedico {
				risco++
			}
		}
	}

	fmt.Println(risco)
}
