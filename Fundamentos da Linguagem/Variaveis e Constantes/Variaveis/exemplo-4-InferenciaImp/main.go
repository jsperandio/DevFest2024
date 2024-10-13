package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Ate agora declaramos variaveis de forma explicita
	var valor int = 10
	fmt.Println("valor:", valor)

	// mas agora vamos inferir o tipo de variavel
	// o tipo de variavel e o valor atribuido para a variavel
	valorImplicito := 10
	fmt.Println("valor implicito:", valorImplicito)

	// o mesmo vale para strings
	textoImplicito := "implicito"
	fmt.Println("texto implicito:", textoImplicito)

	// cuidado com floats
	// se nao definirmos a casa decimal o go interpreta como um int
	floatImplicito := 10
	fmt.Println("float implicito:", floatImplicito, "tipo:", reflect.TypeOf(floatImplicito))
	// o correto seria
	floatImplicitoCorreto := 10.0
	fmt.Println("float implicito correto:", floatImplicitoCorreto, "tipo:", reflect.TypeOf(floatImplicitoCorreto))

	// tambem podemos usar com a keyword var
	// mas usar acaba sendo desvantajoso, 3 letras quando temos o walrus operator
	// acaba nem sendo muito utilizado
	var varImplicita = "varImplicita"
	fmt.Println("var implicita:", varImplicita)
}
