package main

import "fmt"

func main() {

	// Operadores logicos
	// https://go.dev/ref/spec#Logical_operators

	verdadeiro := true
	falso := false

	fmt.Println("AND = verdadeiro && falso:", verdadeiro && falso)
	fmt.Println("OR = verdadeiro || falso:", verdadeiro || falso)
	fmt.Println("NOT = !verdadeiro:", !verdadeiro)
}
