package main

import "fmt"

func main() {
	var h1, m1, s1, h2, m2, s2 int

	fmt.Scan(&h1, &m1, &s1)
	fmt.Scan(&h2, &m2, &s2)

	t1 := h1*3600 + m1*60 + s1
	t2 := h2*3600 + m2*60 + s2

	if t2 <= t1 {
		t2 += 24 * 3600
	}

	dif := t2 - t1

	h := dif / 3600
	m := (dif % 3600) / 60
	s := dif % 60

	fmt.Printf("%02d %02d %02d\n", h, m, s)
}