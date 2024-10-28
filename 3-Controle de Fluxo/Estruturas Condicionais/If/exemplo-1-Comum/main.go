package main

import "fmt"

func main() {
	// exemplo comum de if
	// e algumas formas de comparação
	condicao := true
	if condicao {
		fmt.Println("verdadeiro")
	}

	if (10 + 5) >= 15 {
		fmt.Println("If comparação maior ou igual")
	}

	if condicao && (10+5) >= 15 {
		fmt.Println("If comparação AND")
	}

	if !condicao || (10+5) < 15 {
		fmt.Println("If comparação OR")
	}

}
