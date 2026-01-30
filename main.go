package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Pizza struct {
	ID int `json:"id"`
	Nome string `json:"nome"`
	Preco float64 `json:"preco"`
}

func main() {
	router := gin.Default()

	router.GET("/pizzas", getPizzas)

	router.Run()
}

func getPizzas(c *gin.Context) {
	var pizzas = []Pizza{
		{ID: 1, Nome: "Toscana", Preco: 45.00},
		{ID: 2, Nome: "Marguerita", Preco: 35.00},
		{ID: 3, Nome: "Calabresa", Preco: 40.00},
	}
	fmt.Println(pizzas)

	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}