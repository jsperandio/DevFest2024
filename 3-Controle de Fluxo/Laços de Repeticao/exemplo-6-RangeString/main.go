package main

import "fmt"

func main() {
	// exemplo de um for range para strings
	msg := "Ola Mundo!"

	// para strings o valor do char é um rune
	// o que é um rune? https://go.dev/ref/spec#Numeric_types
	// rune é um alias de int32
	for i, letra := range msg {
		fmt.Println(i, string(letra))
	}

}
