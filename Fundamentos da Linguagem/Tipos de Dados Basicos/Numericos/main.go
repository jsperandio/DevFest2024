package main

import "fmt"

func exemploInteirosEUnsigned() {
	fmt.Println("exemploInteirosEUnsigned:")
	// Tipo int
	var a int = 10
	// Tipo int8 valores de -128 a 127
	var b int8 = 127
	// Tipo int16 valores de -32768 a 32767
	var c int16 = 32767
	// Tipo int32 valores de -2147483648 a 2147483647
	var d int32 = 2147483647
	// Tipo int64 valores de -9223372036854775808 a 9223372036854775807
	var e int64 = 9223372036854775807
	// Tipo uint valores de 0 a 4294967295
	var f uint = 4294967295
	// Tipo uint8 valores de 0 a 255
	var g uint8 = 255
	// Tipo uint16 valores de 0 a 65535
	var h uint16 = 65535
	// Tipo uint32 valores de 0 a 4294967295
	var i uint32 = 4294967295
	// Tipo uint64 valores de 0 a 18446744073709551615
	var j uint64 = 18446744073709551615

	fmt.Println(a, b, c, d, e, f, g, h, i, j)
}

func exemploFloat() {
	fmt.Println("exemploFloat:")
	// Tipo float32 valores de -3.40282346638528859811704183484516925440e+38 a 3.40282346638528859811704183484516925440e+38
	var a float32 = 3.40282346638528859811704183484516925440e+38
	// Tipo float64 valores de -1.797693134862315708145274237317043567981e+308 a 1.797693134862315708145274237317043567981e+308
	var b float64 = 1.797693134862315708145274237317043567981e+308

	fmt.Println(a, b)
}

func exemploComplex() {
	fmt.Println("exemploComplex:")
	// Tipo complex64 valores de -3.40282346638528859811704183484516925440e+38 a 3.40282346638528859811704183484516925440e+38
	var a complex64 = 3.40282346638528859811704183484516925440e+38
	// Tipo complex128 valores de -1.797693134862315708145274237317043567981e+308 a 1.797693134862315708145274237317043567981e+308
	var b complex128 = 1.797693134862315708145274237317043567981e+308

	fmt.Println(a, b)
}

func main() {
	// 90% de todos os usos serao de inteiros.
	// e provavelmente só [int,int64]
	exemploInteirosEUnsigned()
	// 90% de todos os usus serao float64
	exemploFloat()
	// nicho muito raro para uso
	exemploComplex()
}
