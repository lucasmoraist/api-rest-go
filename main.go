package main

import "fmt"

func main() {
	// var nomePizzaria string = "Pizzaria Go"
	// fmt.Println(nomePizzaria)

	// Dessa maneira, o Go infere o tipo da variável automaticamente
	nomePizzaria := "Pizzaria Go"
	instagram, telefone := "@pizzariago", 12345678
	fmt.Println(nomePizzaria, instagram, telefone)
}
