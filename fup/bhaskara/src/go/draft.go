package main
import (
	"fmt"
	"math"
)
func main() {
	var a, b, c float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	delta := b*b - 4*a*c

	if delta > 0 {
		x1 := (b*(-1) + math.Sqrt(delta)) / (2 * a)
		x2 := (b*(-1) - math.Sqrt(delta)) / (2 * a)
           fmt.Printf("%.2f\n%.2f\n", x1, x2)
	} else if delta == 0 {
        if b==0 {
            fmt.Printf("%.2f\n", b)
            return 
        }
		x := b*(-1) / (2 * a)
		fmt.Printf("%.2f\n", x)
	} else {
		fmt.Println("nao ha raiz real")
	}
}

