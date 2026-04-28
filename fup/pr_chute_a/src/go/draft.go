 package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {  
	rand.Seed(time.Now().UnixNano())
	limiteInferior := 1
	limiteSuperior := 100
	numeroSecreto := rand.Intn(limiteSuperior-limiteInferior-1) + limiteInferior + 1
	var chute int
	for {
		if limiteSuperior-limiteInferior <= 2 {
			fmt.Println("Você perdeu! Não há mais números possíveis.")
			break
		}
		fmt.Printf("Digite um número entre %d e %d: ", limiteInferior, limiteSuperior)
		fmt.Scan(&chute)
		if chute <= limiteInferior || chute >= limiteSuperior {
			fmt.Println("Chute inválido! Deve estar dentro do intervalo aberto.")
			continue
		}
		if chute == numeroSecreto {
			fmt.Println("Parabéns! Você acertou!")
			break
		} else if chute > numeroSecreto {
			limiteSuperior = chute
		} else {
			limiteInferior = chute
		}
	}
}

