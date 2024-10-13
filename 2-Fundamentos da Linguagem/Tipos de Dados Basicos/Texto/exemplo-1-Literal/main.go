package main

import "fmt"

func main() {

	// Declaracao literal de uma string
	// preservam o formato, incluindo quebras de linha
	var msg string
	msg = `Hello,
World!
	`
	fmt.Println(msg)
}
