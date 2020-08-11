// Declarando variavel com package scope.
// Toda variavel 

package main

import "fmt"

var nome string
var idade int
var peso float64
var solteiro bool

func main() {
   fmt.Printf("%v, %T\n", nome, nome)
   fmt.Printf("%v, %T\n", idade, idade)
   fmt.Printf("%v, %T\n", peso, peso)
   fmt.Printf("%v, %T\n", solteiro, solteiro)
   
   nome := "Douglas"
   if !solteiro {
	   fmt.Println(nome, "esta compromissado atualmente!")
   }

}

