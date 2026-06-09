package main

import (
	"fmt"
)

func main() {
	var h, m, s int
	_, err := fmt.Scan(&h, &m, &s)
	if err != nil {
		return
	}

	s++

	if s == 60 {
		s = 0
		m++
		if m == 60 {
			m = 0
			h++
			if h == 24 {
				h = 0
			}
		}
	}

	fmt.Printf("%02d %02d %02d\n", h, m, s)
}