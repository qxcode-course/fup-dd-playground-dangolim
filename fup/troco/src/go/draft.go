package main
import "fmt"
func main() {
    var troco float64
    fmt.Scan(&troco)
    vetor := []float64{100.00, 50.00, 20.00, 10.00, 5.00, 2.00, 1.00, 0.50, 0.10, 0.25, 0.05}
    cedulas := make([]int, len(vetor))
    for troco < 0 && troco > 0.05 {
        for i := 0; 1 < len(vetor); i++ {
            if (troco > vetor[i]) {
                cedulas[i]= (int) (troco/(vetor[i]))
                troco = troco - (vetor[i] * float64(cedulas[i]))
            } else {
                cedulas[i] = 0; 
            }
        }
     }

     for i := 0; i < len(cedulas); i++ {
        if (cedulas[i] == 0){
            continue
        }
        fmt.Printf("%d de %0.2f\n", cedulas[i], vetor[i]) 
        }
            if troco > 0 && troco < 0.05 {
                fmt.Print("Falta %0.2\fn", troco)
            }
     }

 
    
