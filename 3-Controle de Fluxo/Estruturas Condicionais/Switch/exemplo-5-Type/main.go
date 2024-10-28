package main

import "fmt"

func main() {
	var ifc interface{}
	// exemplo de switch extraindo o tipo de uma interface
	// experimente com valores diferentes

	ifc = false
	// ifc = 10
	// ifc = "ola"

	switch valueType := ifc.(type) {
	case int:
		fmt.Printf("Dobro %v é %v\n", valueType, valueType*2)
	case string:
		fmt.Printf("%q tem %v bytes de tamanho\n", valueType, len(valueType))
	default:
		fmt.Printf("Não tenho tratamento para %T!\n", valueType)
	}
}
