package main

import "fmt"

func main() {
	// Slice é uma estrutura flexível e dinâmica que permite trabalhar com uma sequência de elementos.
	// para declarar um slice, você pode usar os colchetes [] sem especificar o tamanho,
	// diferente de um array.
	var dArr []int
	dArr = append(dArr, 1) // Adicionando um elemento
	dArr = append(dArr, 2) // Adicionando um elemento
	dArr = append(dArr, 3) // Adicionando um elemento
	fmt.Println(dArr)

}

// Nerd only
// internamente, um slice ""é"" um array
// caso ainda tenha curiosidade de como funciona por baixo dos panos:
// https://dev.to/jpoly1219/how-slices-work-in-go-47nc
// https://go.dev/ref/spec#Slice_types
// e se ainda tiver mais interesse, uma leitura mais tecnica:
// https://victoriametrics.com/blog/go-slice/
