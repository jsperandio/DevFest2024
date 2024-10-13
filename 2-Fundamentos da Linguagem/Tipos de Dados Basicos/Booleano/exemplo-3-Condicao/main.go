package main

import "fmt"

func main() {
	var verdadeiro bool
	var falso bool

	verdadeiro = 1 == 1
	falso = 1 >= 2

	fmt.Println(verdadeiro)
	fmt.Println(falso)
}
