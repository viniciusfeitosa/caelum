package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	totalDeVidas = 3
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
	param := os.Args[1]

	vidas, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Digite a letra para a palavra secreta")
	fmt.Println("Você tem", vidas, "vidas")

	var chute string
	for vidas > 0 {
		fmt.Println(strings.Join(acertos, " "))

		if palavraSecreta == strings.Join(acertos, "") {
			fmt.Println("Fim de Jogo, você venceu")
			break
		}

		fmt.Scanln(&chute)
		if strings.ContainsAny(palavraSecreta, chute) {
			fmt.Println("Você acertou")
			geraSequencia(palavraSecreta, chute)
			continue
		}
		fmt.Println("Você errou")
		vidas--
	}

	if vidas == 0 {
		fmt.Println("Fim de Jogo, você perdeu")
	}
}

func geraSequencia(palavraSecreta, chute string) {
	for indice, letra := range palavraSecreta {
		if string(letra) == chute {
			acertos[indice] = chute
		}
	}
}
