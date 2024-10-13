package main

import "fmt"

func main() {
	// Um mapa é um grupo não ordenado de elementos de um tipo,
	// indexado por um conjunto de chaves unicas de outro tipo.
	// https://go.dev/ref/spec#Map_types
	// Pense como um dicionario, uma coleção de pares chave-valor
	// O valor de um mapa não inicializado é NULO.
	var ages map[string]int

	// Para inicializar um mapa, basta usar a sintaxe make(t Type, size ...int)
	// o make aloca um espaco na memoria : https://pkg.go.dev/builtin#make
	ages = make(map[string]int)

	// Para adicionar um par chave-valor, basta usar a sintaxe
	// mapa[chave] = valor
	ages["João"] = 30
	ages["Maria"] = 25
	ages["Pedro"] = 40
	// caso tente adicionar um par chave-valor que ja existe,
	// o valor de chave existente será atualizado
	ages["Maria"] = 0

	// Para remover um par chave-valor, basta usar a sintaxe
	// delete(mapa, chave)
	delete(ages, "Pedro")

	fmt.Println("Map:", ages)
}
