package main

import "fmt"

// Variaveis globais podem ser acessadas por todas as funções
// elas pertencem ao pacote e existem ate o final do programa
// cuidado ao utilizalas pois elas nao sao coletadas pelo garbage collector
// mesmo que não tenham sido utilizadas
var euSouGlobal string = "global"

func main() {
	// Variaveis locais podem ser acessadas apenas dentro da funcao
	// ela pertence apenas ao escopo da funcao
	// e sao coletadas pelo garbage collector quando *encerra a funcao
	var euSouLocal string = "local"

	fmt.Println(euSouGlobal)
	fmt.Println(euSouLocal)

	if euSouDeBloco := "bloco"; euSouDeBloco == "bloco" {
		fmt.Println(euSouDeBloco)
	}
	// nao funciona, pois euSouDeBloco existe somente dentro do bloco if
	// fmt.Println(euSouDeBloco)

}
