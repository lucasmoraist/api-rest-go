package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pizzaria/models"
)

func main() {
	router := gin.Default()

	router.GET("/pizzas", getPizza)

	router.Run()
}

func getPizza(c *gin.Context) {
	pizzas := []models.Pizza{
		{ID: 1, Nome: "Margherita", Preco: 25.0},
		{ID: 2, Nome: "Pepperoni", Preco: 30.0},
		{ID: 3, Nome: "Quatro Queijos", Preco: 35.0},
	}

	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
	fmt.Println("Pizzas retornadas com sucesso")
}