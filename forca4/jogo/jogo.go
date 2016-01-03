package jogo

import (
	"fmt"
	"strings"
)

// Jogo é a estrutura que contem os dados para um novo jogo
type Jogo struct {
	acertos        []string
	palavraSecreta string
	vidas          int64
}

const (
	fimDeJogoSucesso = "Fim de Jogo, você venceu!!!"
	fimDeJogoFalha   = "Fim de Jogo, você perdeu."
	acertou          = "Você acertou!!!"
	errou            = "Você errou."
)

// NewJogo gera uma nova instancia de Jogo
func NewJogo(palavraSecreta string, vidas int64) *Jogo {
	acertos := make([]string, len(palavraSecreta))
	for i := range acertos {
		acertos[i] = "_"
	}
	return &Jogo{acertos: acertos, palavraSecreta: palavraSecreta, vidas: vidas}
}

// Play é a função inicializadora do jogo de forca
func (j *Jogo) Play() {

	if chegouFim, msg := j.verificaFimJogo(); chegouFim {
		fmt.Println(msg)
		return
	}

	var chute string
	fmt.Scanln(&chute)
	if strings.Contains(j.palavraSecreta, chute) {
		fmt.Println(acertou)
		j.geraSequencia(j.palavraSecreta, chute)
	} else {
		fmt.Println(errou)
		j.vidas--
	}
	j.Play()
}

func (j *Jogo) verificaFimJogo() (bool, string) {
	if j.vidas < 1 {
		return true, fimDeJogoFalha
	}
	fmt.Println(strings.Join(j.acertos, " "))
	if j.palavraSecreta == strings.Join(j.acertos, "") {
		return true, fimDeJogoSucesso
	}
	return false, ""
}

func (j *Jogo) geraSequencia(palavraSecreta, chute string) {
	for indice, letra := range palavraSecreta {
		if string(letra) == chute {
			j.acertos[indice] = chute
		}
	}
}
