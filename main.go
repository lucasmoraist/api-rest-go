package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pizzaria/models"
	"strconv"
)

var pizzas = []models.Pizza{
	{ID: 1, Nome: "Margherita", Preco: 25.0},
	{ID: 2, Nome: "Pepperoni", Preco: 30.0},
	{ID: 3, Nome: "Quatro Queijos", Preco: 35.0},
}

func main() {
	router := gin.Default()

	router.GET("/pizzas", getPizza)
	router.POST("/pizzas", postPizzas)
	router.GET("/pizzas/:id", getPizzasById)

	router.Run()
}

func getPizza(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
	fmt.Println("Pizzas retornadas com sucesso")
}

func postPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pizzas = append(pizzas, newPizza)
}

func getPizzasById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
 
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	for _, p := range pizzas {
		if p.ID == id {
			c.JSON(200, p)
			return
		}
	}

	c.JSON(404, gin.H{"error": "Pizza não encontrada"})

}
