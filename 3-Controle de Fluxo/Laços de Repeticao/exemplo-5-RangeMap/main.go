package main

import "fmt"

func main() {
	// exemplo de um for range para mapas
	dias := map[int]string{
		2: "segunda",
		3: "terça",
		4: "quarta",
		5: "quinta",
		6: "sexta",
		7: "sabado",
		1: "domingo",
	}

	// com o mapa, o range retorna o key e o value
	for key, value := range dias {
		fmt.Println(key, value)
	}

}
