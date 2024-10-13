package main

import "fmt"

// Uma estrutura é uma sequência de elementos nomeados, chamados campos(fields), cada um com um nome e um tipo.
// Dentro de uma estrutura, os nomes dos campos devem ser unicos.
// https://go.dev/ref/spec#Struct_types
type Cidadao struct {
	Nome          string
	Idade         int
	Nacionalidade string
}

func main() {
	var c Cidadao
	c = Cidadao{
		Nome:          "João",
		Idade:         30,
		Nacionalidade: "Brasileiro",
	}

	fmt.Println(c)
	// Voce pode declarar structs dentro de funções
	// Para usos de escopo mais fechado
	type Pessoa struct {
		Nome  string
		Idade int
	}

	var p Pessoa
	p = Pessoa{
		// Para acessar os campos de uma estrutura,
		// basta colocar o nome da estrutura seguido de um ponto e o nome do campo
		Nome:  c.Nome,
		Idade: c.Idade,
	}

	fmt.Println(p)
	fmt.Println(p.Nome)
}
