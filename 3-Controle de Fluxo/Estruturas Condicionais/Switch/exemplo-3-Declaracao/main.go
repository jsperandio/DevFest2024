package main

import "fmt"

func main() {
	// exemplo switch com declaracao de variavel
	// da mesma forma que o if o switch pode ser usado para declarar variaveis no escopo de bloco
	switch dia := "segunda"; dia {
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
