package main

import "fmt"

func exemploDeclaracaoComValor() {
	fmt.Println("exemploDeclaracaoComValor:")

	var verdadeiro bool = true
	var falso bool = false

	fmt.Println(verdadeiro)
	fmt.Println(falso)
}

func exemploDeclaracaoSemValor() {
	fmt.Println("exemploDeclaracaoSemValor:")
	// O valor padrao de um bool é false
	var verdadeiro bool
	var falso bool

	// O valor de falso pode ser alterado
	verdadeiro = true

	fmt.Println(verdadeiro)
	fmt.Println(falso)
}

func main() {
	exemploDeclaracaoComValor()
	exemploDeclaracaoSemValor()
}
