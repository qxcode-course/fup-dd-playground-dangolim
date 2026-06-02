package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    led := map[rune]int{
        '0': 6, '1': 2, '2': 5, '3': 5, '4': 4, '5':5, '6':6, '7':3, '8':7, '9': 6,
    }
    for ; n > 0; n-- {
        var s string 
        fmt.Scan(&s)
        total := 0
        for _, c := range s {
            total += led[c]
        }
        fmt.Println(total, "leds")
    }
}