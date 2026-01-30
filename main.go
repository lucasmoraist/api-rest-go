package main

import "fmt"

type Pizza struct {
	id int
	nome string
	preco float64
}

func main() {
	var pizzas = []Pizza{
		{id: 1, nome: "Toscana", preco: 45.00},
		{id: 2, nome: "Marguerita", preco: 35.00},
		{id: 3, nome: "Calabresa", preco: 40.00},
	}

	fmt.Println("Cardápio:", pizzas)
}
