package main
import "fmt"
func main() {
    var totalsegundos int 
    var horas, minutos, segundos int
    fmt.Scan(&totalsegundos)
    horas = totalsegundos/3600
    resto:= totalsegundos % 3600
    minutos = resto /60
    segundos= resto % 60
    fmt.Printf( "%d:%d:%d\n", horas, minutos, segundos)
}
