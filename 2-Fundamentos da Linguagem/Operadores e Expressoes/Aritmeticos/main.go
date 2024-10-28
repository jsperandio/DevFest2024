package main

import "fmt"

func main() {
	// Semelhante a muitas linguagens de programação
	// Go pode fazer operações aritmeticas
	// https://go.dev/ref/spec#Operators
	// https://go.dev/ref/spec#Arithmetic_operators

	var (
		x int = 10
		y int = 10
	)
	fmt.Printf("Dado os valores X = %d e Y = %d\n", x, y)

	// Funciona para inteiros e floats e pode ser usado com strings
	rst := x + y
	txt := "So" + "ma"
	fmt.Println(txt, rst)

	rst = x - y
	fmt.Println("Subtração:", rst)

	rst = x * y
	fmt.Println("Multiplicação:", rst)

	rst = x / y
	fmt.Println("Divisão:", rst)

	rst = x % y
	fmt.Println("Resto:", rst)

}
