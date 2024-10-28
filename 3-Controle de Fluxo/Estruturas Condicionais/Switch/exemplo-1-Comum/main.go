package main

import "fmt"

func main() {
	// exemplo comum de switch

	dia := "segunda"
	switch dia {
	case "segunda":
		fmt.Println("Hoje é segunda-feira")
	case "terça":
		fmt.Println("Hoje é terça-feira")
	case "quarta":
		fmt.Println("Hoje é quarta-feira")
	default:
		fmt.Println("Hoje é um dia da semana desconhecido")
	}
}
