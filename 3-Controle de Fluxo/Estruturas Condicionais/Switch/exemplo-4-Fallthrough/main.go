package main

import "fmt"

func main() {
	// exemplo de switch com fallthrough
	// o switch vai executar o case que for verdadeiro, e com o fallthrough ele vai executar o proximo case
	// mesmo que ele seja falso

	switch dia := "segunda"; dia {
	case "segunda":
		fmt.Println("Hoje é segunda-feira")
		fallthrough
	case "terça":
		fmt.Println("Hoje é terça-feira")
	case "quarta":
		fmt.Println("Hoje é quarta-feira")
	default:
		fmt.Println("Hoje é um dia da semana desconhecido")
	}
}
