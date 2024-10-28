package main

import "fmt"

func main() {
	// exemplo de um for infinito
	// para cancelar a execução basta dar ctrl+c no terminal que executou o go run
	i := 0
	for {
		i++
		fmt.Println("Infinito, iteração: ", i)
	}
}
