package jogo

import (
	"fmt"
	"strings"
)

const (
	fimDeJogoSucesso = "Fim de Jogo, você venceu!!!"
	fimDeJogoFalha   = "Fim de Jogo, você perdeu."
	acertou          = "Você acertou!!!"
	errou            = "Você errou."
)

var acertos []string
var palavraSecreta = "MELANCIA"

func init() {
	acertos = make([]string, len(palavraSecreta))

	for i := range acertos {
		acertos[i] = "_"
	}
}

// Play é a função inicializadora do jogo de forca
func Play(vidas int64) {
	if vidas < 1 {
		fmt.Println(fimDeJogoFalha)
		return
	}

	fmt.Println(strings.Join(acertos, " "))
	if palavraSecreta == strings.Join(acertos, "") {
		fmt.Println(fimDeJogoSucesso)
		return
	}

	var chute string
	fmt.Scanln(&chute)
	if strings.Contains(palavraSecreta, chute) {
		fmt.Println(acertou)
		geraSequencia(palavraSecreta, chute)
	} else {
		fmt.Println(errou)
		vidas--
	}
	Play(vidas)
}

func geraSequencia(palavraSecreta, chute string) {
	for indice, letra := range palavraSecreta {
		if string(letra) == chute {
			acertos[indice] = chute
		}
	}
}
