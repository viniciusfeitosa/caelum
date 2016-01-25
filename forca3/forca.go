package main

import (
	"fmt"
	"github.com/viniciusfeitosa/caelum/forca3/jogo"
	"log"
	"os"
	"strconv"
)

func main() {
	param := os.Args[1]

	vidas, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Digite a letra para a palavra secreta")
	fmt.Println("Você tem", vidas, "vidas")
	jogo.Play(vidas)
}
