package main

import "fmt"

func main() {
	// Ate agora nao temos a explicacao de uma funcao
	// porem e possivel criar uma variavel tipo funcao
	// https://go.dev/ref/spec#Function_types
	var somar func(x int, y int) int
	var resp int = 0

	// o nome dos paramentros não importam desde que sejam do mesmo tipo
	somar = func(valorA int, valorB int) int {
		return valorA + valorB
	}

	resp = somar(10, 10)
	fmt.Println(resp) // Saída: 15
}
