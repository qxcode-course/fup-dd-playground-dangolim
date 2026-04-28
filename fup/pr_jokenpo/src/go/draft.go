package main
import (
	"fmt"
	"math/rand"
	"time"
)
const (
	Pedra   = 1
	Papel   = 2
	Tesoura = 3
)
func escolhaParaString(escolha int) string {
	switch escolha {
	case Pedra:
		return "Pedra"
	case Papel:
		return "Papel"
	case Tesoura:
		return "Tesoura"
	default:
		return "Inválido"
	}
}
func determinarVencedor(jogador, computador int) int {
	if jogador == computador {
		return 0 // empate
	}
	if (jogador == Pedra && computador == Tesoura) ||
		(jogador == Tesoura && computador == Papel) ||
		(jogador == Papel && computador == Pedra) {
		return 1 // jogador vence
	}

	return 2 // computador vence
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var jogarNovamente string

	for {
		pontosJogador := 0
		pontosComputador := 0

		fmt.Println("=== Jokenpô (Pedra, Papel e Tesoura) ===")

		for round := 1; round <= 5; round++ {
			var escolhaJogador int

			fmt.Printf("\nRound %d de 5\n", round)
			fmt.Println("Escolha: 1-Pedra | 2-Papel | 3-Tesoura")
			fmt.Print("Sua jogada: ")
			fmt.Scan(&escolhaJogador)

			if escolhaJogador < 1 || escolhaJogador > 3 {
				fmt.Println("Jogada inválida! Tente novamente.")
				round--
				continue
			}

			escolhaComputador := rand.Intn(3) + 1

			fmt.Printf("Você escolheu: %s\n", escolhaParaString(escolhaJogador))
			fmt.Printf("Computador escolheu: %s\n", escolhaParaString(escolhaComputador))

			resultado := determinarVencedor(escolhaJogador, escolhaComputador)

			switch resultado {
			case 0:
				fmt.Println("Empate!")
			case 1:
				fmt.Println("Você venceu o round!")
				pontosJogador++
			case 2:
				fmt.Println("Computador venceu o round!")
				pontosComputador++
			}

			fmt.Printf("Placar -> Você: %d | Computador: %d\n", pontosJogador, pontosComputador)
		}

		fmt.Println("\n=== Resultado Final ===")
		fmt.Printf("Você: %d | Computador: %d\n", pontosJogador, pontosComputador)

		if pontosJogador > pontosComputador {
			fmt.Println("Você venceu o jogo!")
		} else if pontosComputador > pontosJogador {
			fmt.Println("Computador venceu o jogo!")
		} else {
			fmt.Println("O jogo terminou em empate!")
		}

		fmt.Print("\nDeseja jogar novamente? (s/n): ")
		fmt.Scan(&jogarNovamente)

		if jogarNovamente != "s" && jogarNovamente != "S" {
			fmt.Println("Obrigado por jogar!")
			break
		}
	}
}