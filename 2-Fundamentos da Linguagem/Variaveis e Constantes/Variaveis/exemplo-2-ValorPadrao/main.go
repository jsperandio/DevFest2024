package main

import "fmt"

func main() {
	// Todos os tipos de dados em Go possuem um valor padrão
	// https://go.dev/ref/spec#The_zero_value

	var inteiro int       // Valor padrao 0
	var flutuante float64 // Valor padrao 0.0
	var str string        // Valor padrao e uma string vazia ""
	var check bool        // Valor padrao false
	var ptr *int          // Valor padrao (nil)
	var slc []int         // Valor padrao (nil)
	var m map[string]int  // Valor padrao (nil)

	type Example struct {
		X     int
		Texto string
	}

	var e Example // Valor padrao sera o valor zero de cada tipo dos campos

	fmt.Println("inteiro:", inteiro)
	fmt.Println("flutuante:", flutuante)
	fmt.Println("string:", str)
	fmt.Println("bool:", check)
	fmt.Println("ponteiro:", ptr)
	fmt.Println("slice:", slc, "tamanho:", len(slc), "é nulo? ", slc == nil)
	fmt.Println("map:", m, "tamanho:", len(m), "é nulo? ", m == nil)
	fmt.Println("struct:", e)
	// ué , porque o valor de um ponteiro e nil e exibe nil
	// enquanto o valor de um slice,map é nil e exibe []?
	// vamos checar
	fmt.Println("======")
	slc = make([]int, 0)
	m = make(map[string]int)
	fmt.Println("slice:", slc, "tamanho:", len(slc), "é nulo? ", slc == nil)
	fmt.Println("map:", m, "tamanho:", len(m), "é nulo? ", m == nil)
	// Existe diferenca entre nil(nulo) e vazio,
	// slice vazio tem uma alocação de memória, enquanto o slice nil não.
	// o mesmo regra se aplica para o map
}
