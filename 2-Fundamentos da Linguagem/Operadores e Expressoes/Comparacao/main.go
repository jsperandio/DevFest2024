package main

import "fmt"

func main() {

	// Operadores de Comparação
	// https://go.dev/ref/spec#Comparison_operators

	var (
		x int = 10
		y int = 10
	)
	fmt.Printf("Dado os valores X = %d e Y = %d\n", x, y)

	// ==
	fmt.Println("Igualdade:", x == y)
	// !=
	fmt.Println("Desigualdade:", x != y)
	// >
	fmt.Println("Maior que:", x > y)
	// <
	fmt.Println("Menor que:", x < y)
	// >=
	fmt.Println("Maior ou igual:", x >= y)
	// <=
	fmt.Println("Menor ou igual:", x <= y)

	// Comparando strings
	fmt.Println("Comparando strings:", "abc" == "abc")
	fmt.Println("Comparando strings:", "abc" != "def")

	// Comparando ponteiros (address)
	// os simbolos *(asterisco) e &(ampersand)
	// fazem acesso ao valor do ponteiro e
	// a referência da memoria respectivamente
	// https://go.dev/ref/spec#Address_operators
	var a, b *int

	a = &x
	b = &y

	fmt.Println("Comparando igualdade ponteiros:", a == b)
	fmt.Println("Comparando diferenca ponteiros:", a != b)
	fmt.Println("Comparando igualdade valor ponteiros:", *a == *b)

}
