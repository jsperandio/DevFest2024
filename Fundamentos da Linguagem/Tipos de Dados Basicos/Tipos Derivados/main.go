package main

import "fmt"

type Cidadao struct {
	Nome          string
	Idade         int
	Nacionalidade string
}

func examploArray() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)
}

func examploSlice() {
	dArr := []int{1, 2, 3, 4, 5} // Slice dinâmico
	dArr = append(dArr, 6)       // Adicionando um elemento
	fmt.Println("Slice:", dArr)
}

// Função para exemplificar Structs
func examploStruct() {
	var c Cidadao
	c = Cidadao{
		Nome:          "João",
		Idade:         30,
		Nacionalidade: "Brasileiro",
	}

	fmt.Println("Struct Cidadao:", c)
	// Voce pode declarar structs dentro de funções
	// Para usos de escopo mais fechado
	type Pessoa struct {
		Nome  string
		Idade int
	}

	var p Pessoa
	p = Pessoa{
		Nome:  c.Nome,
		Idade: c.Idade,
	}

	fmt.Println("Struct Pessoa:", p)
}

func examploMap() {
	var ages map[string]int

	ages = make(map[string]int)

	ages["João"] = 30
	ages["Maria"] = 25

	fmt.Println("Map:", ages)
}

func examploPointer() {
	var val int
	val = 42

	// Semelhante a linguagem C
	// Para definir um ponteiro usamos o *
	var prt *int
	// & representa o endereço de uma variável na memória
	prt = &val
	// Caso usemos o ponteiro direto , mostrara o valor do endereco de memoria
	fmt.Println("Pointer:", prt)
	fmt.Println("Pointer Value:", *prt)
}

// Função para exemplificar Func
func examploFunc() {
	var somar func(x int, y int) int
	var resp int = 0

	somar = func(valorA int, valorB int) int {
		return valorA + valorB
	}

	resp = somar(10, 10)
	fmt.Println("Resultado variavel tipo func somar:", resp) // Saída: 15
}

func main() {
	examploArray()
	examploSlice()
	examploStruct()
	examploMap()
	examploPointer()
	examploFunc()
}
