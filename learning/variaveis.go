package main

import (
	"fmt"
)

var nome string = "Caio"

var (
	nome2 string = "Nome"
	nome3 string = "Nome3"
	nome4 string = "Nome4"
)

// First declaration
var i int = 2
var I int = 3

func main() {
	// Ways to create a var
	// Second declaration
	var i int
	i = 1
	fmt.Println(i)

	var a int = 2
	fmt.Println(a)

	b := 10
	fmt.Println(b)

	fmt.Println(nome)
	fmt.Println(nome2)
	fmt.Println(nome3)
	fmt.Println(nome4)

	fmt.Println(I)

}
