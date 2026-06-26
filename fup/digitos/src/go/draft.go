package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if b == 0 {
		if a == 0 {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
		return
	}

	cont := 0
	for b > 0 {
		if b%10 == a {
			cont++
		}
		b /= 10
	}

	fmt.Println(cont)
}