package main

import "fmt"

func main() {
	// if com else if
	x := 5
	if x > 10 {
		fmt.Println("x é maior que 10")
	} else if x > 5 {
		fmt.Println("x é maior que 5, mas menor ou igual a 10")
	} else {
		fmt.Println("x é 5 ou menor")
	}
}
