package main

import "fmt"

const nacionalidadePadrao = "Brasileira"

func main() {
	// Os tipos estao sendo inferidos pelo valor atribuido
	const maiorIdadePadrao = 18

	fmt.Printf(" A maior idade da nacionalidade padrao(%s) é %d anos \n", nacionalidadePadrao, maiorIdadePadrao)
}
