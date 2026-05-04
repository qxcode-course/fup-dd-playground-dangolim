package main
import "fmt"
func main() {
    var c, bananas, goiabas, mangas int
    fmt.Scan(&c, &bananas, &goiabas, &mangas)
    total := bananas + goiabas + mangas
    minutos := total / c
    if total%c !=0 {
        minutos++
    }
    fmt.Println(minutos)
}
