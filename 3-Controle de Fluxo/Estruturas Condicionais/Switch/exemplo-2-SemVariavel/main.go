package main

import "fmt"

func main() {
	// exemplo comum de switch
	// se não especificar variavel realizamos condicoes separadas pelos cases
	// que podem comparar varias variaveis
	dia := "segunda"
	hora := 14
	switch {
	case dia == "segunda" && hora < 12:
		fmt.Println("Hoje é segunda-feira e agora são:", hora, "horas")
	case dia == "segunda" && hora > 12:
		fmt.Println("Hoje é segunda-feira e agora ja passou do meio-dia:", hora, "horas")
	case dia == "terça":
		fmt.Println("Hoje é terça-feira")
	case dia == "quarta":
		fmt.Println("Hoje é quarta-feira")
	default:
		fmt.Println("Hoje é um dia da semana desconhecido")
	}
}
