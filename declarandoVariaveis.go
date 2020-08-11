package main

import "fmt"

var nome string = "Douglas Sumita" //Declaração de variavel de package scope ou Escopo abrangente.

func main() {
	//fmt.Printf("%v, %T\n", nome, nome)
	mostraIdade()
}

func mostraIdade() {
	idade := 25
	fmt.Printf("%v, %T\n", nome, nome)
	fmt.Printf("%v, %T", idade, idade) 
}

