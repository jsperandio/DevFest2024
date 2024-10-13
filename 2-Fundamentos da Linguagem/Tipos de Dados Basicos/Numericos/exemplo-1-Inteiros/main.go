package main

import "fmt"

func main() {
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

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
}
