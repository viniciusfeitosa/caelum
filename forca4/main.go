package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
    "github.com/viniciusfeitosa/caelum/forca4/jogo"
)

func main() {
	qtdVidas := os.Args[1]
    palavraSecreta := os.Args[2]

	vidas, err := strconv.ParseInt(qtdVidas, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Digite a letra para a palavra secreta")
	fmt.Println("Você tem", vidas, "vidas")
    
    novoJogo := jogo.NewJogo(palavraSecreta, vidas)
    novoJogo.Play()	
}