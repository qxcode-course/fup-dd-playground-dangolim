package main

import "fmt"

func main() {
	var n, f, c int
	fmt.Scan(&n)

	for f = 2; n != 1; {
		if n%f == 0 {
			n /= f
			c++
		} else {
			if c > 0 {
				fmt.Println(f, c)
			}
			f++
			c = 0
		}
	}

    

	if c > 0 {
		fmt.Println(f, c)
	}
}
