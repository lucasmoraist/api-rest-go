package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas []models.Pizza

func main() {
	loadPizzas()
	router := gin.Default()

	router.GET("/pizzas", getPizza)
	router.POST("/pizzas", postPizzas)
	router.GET("/pizzas/:id", getPizzasById)

	router.Run()
}

func loadPizzas() {
	file, err := os.Open("data/pizza.json")

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("Erro ao decodificar o arquivo JSON:", err)
	}
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
