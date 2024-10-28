package main

import "fmt"

func main() {
	var x int = 10
	fmt.Println("x externo:", x) // Imprime 10

	if x == 10 {
		// Shadowing
		// a variavel x assume uma nova instancia no bloco
		var x int = 100
		fmt.Println("x interno:", x) // Imprime 5
	}

	fmt.Println("x externo novamente:", x) // Imprime 10, a variável externa não foi alterada

}
