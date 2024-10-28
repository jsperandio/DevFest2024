package main

import "fmt"

func main() {
	// exemplo de um for range
	dias := [7]string{"segunda", "terça", "quarta", "quinta", "sexta", "sabado", "domingo"}

	// como podemos ver, o range retorna o indice e o valor da variável para variaveis do tipo array/slice
	for i, dia := range dias {
		fmt.Println(i, dia)
	}
}
