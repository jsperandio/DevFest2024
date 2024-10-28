package main

import "fmt"

func main() {
	// exemplo de if declaração
	// podemos declarar uma variável dentro do if
	// basta separar por ponto e virgula e depois a expressão de comparação
	if x := 20; x%10 == 0 {
		fmt.Println("x é divisível por 10")
	}
}
