package main
import "fmt"
func main() {
    var c, a int
    fmt.Scan(&c, &a)
    alunosporviagem:= c - 1
    viagens := a / alunosporviagem 
    if a%alunosporviagem !=0 {
        viagens++
    }
    fmt.Println(viagens)
}
