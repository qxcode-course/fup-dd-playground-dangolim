package main
import "fmt"
import "math"
func main() {
    var d, r, x1, y1, x2, y2 float64 
    fmt.Scan(&x1, &y1, &x2, &y2, &r, &d)
    d=math.Pow((x2 - x1), 2) + math.Pow((y2 - y1), 2)
    r= math.Sqrt(d)
    fmt.Printf("%.2f\n", r)
}
