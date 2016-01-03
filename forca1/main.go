package main

import (
	"fmt"
	"strings"
)

var (
	palavraSecreta = "MELANCIA"
	acertos        = make([]string, len(palavraSecreta))
	vidas          = 3
)

func main() {
	for i := range acertos {
		acertos[i] = "_"
	}

	fmt.Println("Digite a letra para a palavra secreta")
	fmt.Println("Você tem", vidas, "vidas")

	var chute string
	for vidas > 0 {
		fmt.Println(strings.Join(acertos, " "))

		if palavraSecreta == strings.Join(acertos, "") {
			fmt.Println("Fim de Jogo, você venceu!!!")
			break
		}

		fmt.Scanln(&chute)
		if strings.Contains(palavraSecreta, chute) {
			fmt.Println("Você acertou!!!")
			for indice, letra := range palavraSecreta {
				if string(letra) == chute {
					acertos[indice] = chute
				}
			}
			continue
		}
		fmt.Println("Você errou.")
		vidas--
	}

	if vidas < 0 {
		fmt.Println("Fim de Jogo, você perdeu.")
	}
}
