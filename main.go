package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da Adivinhação")
	fmt.Println("Um número será sorteado. Tente acertar! O número é um inteiro entre 0 e 100.")
	randomInt := rand.Int64N(101)
	fmt.Println("Número sorteado ", randomInt)
	scanner := bufio.NewScanner(os.Stdin)
	guesses := [10]int64{}

	for i := range guesses {
		fmt.Println("Digite um número:")
		scanner.Scan()
		guessString := strings.TrimSpace(scanner.Text())
		guessInt, err := strconv.ParseInt(guessString, 10, 64)
		if err != nil {
			fmt.Println("O seu chute tem que ser um número inteiro.")
			return
		}
		switch {
		case guessInt < randomInt:
			fmt.Println("Você errou. O número sorteado é maior que o ", guessInt)
		case guessInt > randomInt:
			fmt.Println("Você errou. O número sorteado é menor que o ", guessInt)
		case guessInt == randomInt:
			fmt.Printf(
				"Parabéns você acertou! O número sorteado foi %d.\n"+
					"Você acertou em %d tentativas.\n"+
					"Seus chutes foram: %v\n",
				randomInt, i+1, guesses[:i],
			)
			return
		}

		guesses[i] = guessInt
	}

	fmt.Printf(
		"Você não conseguiu acertar. O número sorteado foi %d.\n"+
			"Seus chutes foram: %v\n",
		randomInt, guesses,
	)
}
