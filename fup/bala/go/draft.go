package main

import (
	"fmt"
	"math"
)
func main() {
    var x1, y1, x2, y2 float64
    var distancia, a, b, c, d float64
    fmt.Scan(&x1, &y1, &x2, &y2)
    c= x2-x1
    d= y2-y1
     a   = math.Pow(c, 2)
     b =  math.Pow(d, 2)
    distancia= math.Sqrt(a+b)
    fmt.Printf("%.2f\n", distancia)
}
