package main

import "fmt"

func main() {
	// Podemos declarar varias variaveis de uma vez
	// com valores iniciais ou nao
	var titulo, descricao string = "Effective Go", "Programming Language"

	titulo = "100 Go Mistakes and How to Avoid Them"
	descricao = "Dodge the most common mistakes made by Go developers"

	fmt.Println("Titulo:", titulo)
	fmt.Println("Descricao:", descricao)

	// As variáveis podem ser declaradas separadamente
	// usando parenteses na keyword var
	// com valores iniciais ou nao
	var (
		autor  string = "Teiva Harsanyi"
		idioma string
	)

	idioma = "EN"

	fmt.Println("Autor:", autor)
	fmt.Println("Idioma:", idioma)

	// Podemos usar inferencia implícita
	preco, paginas := 530.93, 384

	fmt.Println("Preco:", preco)
	fmt.Println("Paginas:", paginas)

}
