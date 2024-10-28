package main

import "fmt"

func main() {
	// Uma atribuição substitui o valor atual armazenado em uma variável
	// por um novo valor especificado por uma expressão.
	// Uma instrução de atribuição pode atribuir um único valor a uma única variável ou
	// vários valores a um número correspondente de variáveis.
	// https://go.dev/ref/spec#Assignment_statements

	var x, y int = 10, 10

	fmt.Printf("Dado os valores X = %d e Y = %d\n", x, y)

	// = - atribuição valor a variável da direita pela esquerda
	x = 20
	fmt.Println("x recebe 20, resultando no valor de X:", x)

	// Operadores de Atribuição compostos
	// realizam operações aritméticas sobre os valores das variáveis da esquerda e da direita
	// e atribuem o resultado a variável da esquerda.

	// += - adiciona o valor da direita ao valor da variável à esquerda e
	// o atribui à variável.
	x += 10
	fmt.Println("x += 10, resultando no valor de X:", x)

	// -= - subtrai e atribui à variável.
	x -= 10
	fmt.Println("x -= 10, resultando no valor de X:", x)

	// *= - multiplica e atribui à variável.
	x *= 10
	fmt.Println("x *= 10, resultando no valor de X:", x)

	// /= - divide...
	x /= 10
	fmt.Println("x /= 10, resultando no valor de X:", x)

	// %= - atribui o resto da divisão ...
	x %= 10
	fmt.Println("x %= 10, resultando no valor de X:", x)

	fmt.Println("============")
	// O operador += pode ser usado para *concatenar strings

	// += - concatena e atribui à variável
	txt := "Hello"
	fmt.Printf("Dado o valor da variável txt = %s\n", txt)
	txt += "World"
	fmt.Println("txt += \"def\", resultando no valor de txt:", txt)

}

// * Embora o operador += seja usado para concatenar strings,
// em situacões em que essa operacao seja realizada dentro de um loop,
// é preferiver usar strings.Builder, Por que ?
//
// Strings em go são imutáveis,
// quando realizamos a operação += não alteramos o valor da variável original
// e sim criamos(alocamos) uma terceira variável com o valor para o resultado da operacao.
//
//
// exemplo da diferença de performance
// https://medium.com/swlh/high-performance-string-building-in-go-golang-3fd99b9ca856
