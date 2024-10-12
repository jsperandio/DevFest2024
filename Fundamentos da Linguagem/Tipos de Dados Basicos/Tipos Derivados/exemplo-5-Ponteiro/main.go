package main

import "fmt"

func main() {
	var valor int
	valor = 999

	// Semelhante a linguagem C
	// Para definir um ponteiro usamos o *
	var prt *int
	// & representa o endereço de uma variável na memória
	prt = &valor
	// Caso usemos o ponteiro direto , mostrara o valor do endereco de memoria
	fmt.Println(prt)
	fmt.Println(*prt)

	// Como modificar o valor de um ponteiro?
	// prt = 0 -> isso não funciona
	// Com o uso do operador *, podemos modificar o valor de um ponteiro
	*prt = 0
	fmt.Println(*prt)
}
