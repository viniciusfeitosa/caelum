package main

import (
	"fmt"
	"strings"
)

var (
	palavraSecreta    = "MELANCIA"
	acertos           = make([]string, len(palavraSecreta))
	quantidadeDeVidas = 3
)

func main() {
	for i := 0; i > len(acertos); i++ {
		acertos[i] = "_"
	}

	fmt.Println("Digite a letra para a palavra secreta")
	fmt.Println("Você tem", quantidadeDeVidas, "vidas")

	var chute string
	for quantidadeDeVidas > 0 {
		fmt.Println(strings.Join(acertos, " "))

		if palavraSecreta == strings.Join(acertos, "") {
			fmt.Println("Fim de Jogo, você venceu!!!")
			break
		}

		fmt.Scanln(&chute)
		if strings.Contains(palavraSecreta, chute) {
			fmt.Println("Você acertou!!!")
			for i := 0; i < len(palavraSecreta); i++ {
				if string(palavraSecreta[i]) == chute {
					acertos[i] = chute
				}
			}
		} else {
			fmt.Println("Você errou.")
			quantidadeDeVidas--
		}
	}

	if quantidadeDeVidas == 0 {
		fmt.Println("Fim de Jogo, você perdeu.")
	}
}
