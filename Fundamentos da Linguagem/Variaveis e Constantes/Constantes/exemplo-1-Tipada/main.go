package main

import "fmt"

const (
	gravidadeTerrestre float32 = 9.8
	pesoMedioHumano    float32 = 70.0
)

func main() {
	// Igual variaveis, constantes compartilham
	// mesma separação de escopo
	const pesoMedio10Humanos float32 = 10 * pesoMedioHumano

	// Fórmula: F = m * g
	forcaGravidade := pesoMedioHumano * gravidadeTerrestre
	forcaGravidade10Humanos := pesoMedio10Humanos * gravidadeTerrestre

	fmt.Printf("A força gravitacional de um humano medio de %.2f kg é: %.2f N\n", pesoMedioHumano, forcaGravidade)
	fmt.Printf("A força gravitacional de 10 humanos pesando %.2f kg é: %.2f N\n", pesoMedio10Humanos, forcaGravidade10Humanos)
}
