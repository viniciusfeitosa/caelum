package main

import (
	"fmt"
	"strings"
)

const (
	digiteUmaLetra       = "Digite a letra para a palavra secreta"
	msgQuantidadeDeVidas = "Você tem %d vidas \n"
	fimDeJogoSucesso     = "Fim de Jogo, você venceu!!!"
	fimDeJogoFalha       = "Fim de Jogo, você perdeu."
	acertou              = "Você acertou!!!"
	errou                = "Você errou."
)

var (
	palavraSecreta    = "MELANCIA"
	acertos           = make([]string, len(palavraSecreta))
	quantidadeDeVidas = 3
)

func init() {
	for i := range acertos {
		acertos[i] = "_"
	}
}

func main() {
	fmt.Println(digiteUmaLetra)
	fmt.Printf(msgQuantidadeDeVidas, quantidadeDeVidas)

	jogo(quantidadeDeVidas)
}

func jogo(quantidadeDeVidas int) {
	if quantidadeDeVidas == 0 {
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
		geraSequenciaDeAcertos(palavraSecreta, chute)
	} else {
		fmt.Println(errou)
		quantidadeDeVidas--
	}

	jogo(quantidadeDeVidas)
}

func geraSequenciaDeAcertos(palavraSecreta, chute string) {
	for indice, letra := range palavraSecreta {
		if string(letra) == chute {
			acertos[indice] = chute
		}
	}
}
