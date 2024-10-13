package main

import "fmt"

func main() {
	// Operadores bitwise
	// Legal ? sim
	// Frequencia de uso ? raro, muito raro
	// Mas ja que possui, vou deixar alguns exemplos, por que? why not ¯\_(ツ)_/¯

	x := 6 // 110
	y := 3 // 011
	fmt.Printf("Valores de X = %d(110) e Y = %d(011)\n", x, y)

	// | OR
	// 110
	// 011
	// 111
	resultado := x | y // 111 = 7
	fmt.Println("Resultado da operacao OR x | y:", resultado, "(111)")

	// & AND
	// 110
	// 011
	// 010
	resultado = x & y // 010 = 2
	fmt.Println("Resultado da operacao AND x & y:", resultado, "(010)")

	// ^ XOR
	// 110
	// 011
	// 101
	resultado = x ^ y // 101 = 5
	fmt.Println("Resultado da operacao XOR x ^ y:", resultado, "(101)")

	// &^ AND NOT
	// 110
	// 011
	// 100
	resultado = x &^ y // 100 = 4
	fmt.Println("Resultado da operacao AND NOT x &^ y:", resultado, "(100)")

	// << SHL desloca para a esquerda
	xSHL := 3             // 011
	resultado = xSHL << 1 // 110, ou seja, 6
	fmt.Println("Resultado da operacao SHL x << 1:", resultado, "(110)")

	// >> SHR desloca para a direita
	xSHR := 6             // 110
	resultado = xSHR >> 1 // 011, ou seja, 3
	fmt.Println("Resultado da operacao SHR x >> 1:", resultado, "(011)")

}
