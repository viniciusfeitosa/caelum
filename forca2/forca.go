package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	fimDeJogoSucesso = "Fim de Jogo, você venceu!!!"
	fimDeJogoFalha   = "Fim de Jogo, você perdeu."
	acertou          = "Você acertou!!!"
	errou            = "Você errou."
)

var (
	palavraSecreta = "MELANCIA"
	acertos        = make([]string, len(palavraSecreta))
)

func init() {
	for i := range acertos {
		acertos[i] = "_"
	}
}

func main() {
	//geraSequenciaInicial()
	param := os.Args[1]

	vidas, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Digite a letra para a palavra secreta")
	fmt.Println("Você tem", vidas, "vidas")

	// game(vidas)
	var chute string
	for vidas > 0 {
		fmt.Println(strings.Join(acertos, " "))

		if palavraSecreta == strings.Join(acertos, "") {
			fmt.Println(fimDeJogoSucesso)
			break
		}

		fmt.Scanln(&chute)
		if strings.ContainsAny(palavraSecreta, chute) {
			fmt.Println(acertou)
			geraSequencia(palavraSecreta, chute)
			continue
		}
		fmt.Println(errou)
		vidas--
	}

	if vidas < 0 {
		fmt.Println(fimDeJogoFalha)
	}

}

func game(vidas int64) {
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
	game(vidas)
}

func geraSequenciaInicial() {
	for i := range acertos {
		acertos[i] = "_"
	}
}

func geraSequencia(palavraSecreta, chute string) {
	for indice, letra := range palavraSecreta {
		if string(letra) == chute {
			acertos[indice] = chute
		}
	}
}
