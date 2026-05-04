package main
import (
	"fmt"
)

func cabe(x, y, h, l int) bool {
	return (x <= h && y <= l) || (x <= l && y <= h)
}

func main() {
	var A, B, C int
	var H, L int

	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)
	fmt.Scan(&H)
	fmt.Scan(&L)

	if cabe(A, B, H, L) || cabe(A, C, H, L) || cabe(B, C, H, L) {
		fmt.Println("S")
	} else {
		fmt.Println("N")
	}
}

