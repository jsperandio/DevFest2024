package main

import "fmt"

func main() {
	// Um Array é uma sequência numerada de elementos de um único tipo,
	// denominado tipo de elemento.
	// O número de elementos é chamado de comprimento do array e nunca é negativo.
	// https://go.dev/ref/spec#Array_types
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
}
