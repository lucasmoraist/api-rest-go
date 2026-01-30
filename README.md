# api-rest-go

## Pré requisitos
- [VSCode](https://code.visualstudio.com/)
- [Go](https://go.dev/)

## Iniciando projeto
1. Executar o comando
```bash
go mod init <nome-do-projeto>
```
2. Criar arquivo main
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```
3. Executar código
```bash
# O ponto executa o arquivo main que está na raiz
go run .

# Também é possível passando o nome do arquivo
go run main.go
```
4. Para buildar
```bash
go build .
```

## Declarando variaveis
- Toda variavel declarada deve ser utilizada, caso contrário o código não compila

### Exemplos
```go
func main() {
	var nomePizzaria string = "Pizzaria Go"
	fmt.Println(nomePizzaria)
}
```
```go
func main() {
	// Dessa maneira, o Go infere o tipo da variável automaticamente
	nomePizzaria := "Pizzaria Go"
	fmt.Println(nomePizzaria)
}
```
```go
func main() {
	nomePizzaria := "Pizzaria Go"
	instagram, telefone := "pizzariago", 12345678
	fmt.Println(nomePizzaria, instagram, telefone)
}
```

## Struct
- Struct é um tipo de dado personalizado que serve para agrupar diferentes campos sob um único nome, semelhante a interfaces do Typescript.
- Quando um campo é criado com a primeira letra minúscula o Go entende ele como privado, quando está com maiúscula ele é público

### Exemplo
```go
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
```

## Gin
### O'que é
O Gin (ou Gin Web Framework) é um framework HTTP escrito em Go (Golang).
Ele é uma das ferramentas mais populares e utilizadas no ecossistema Go para desenvolvimento web.

### Para que ele serve
O principal objetivo do Gin é facilitar e acelerar a criação de APIs RESTful e serviços web backend. Ele serve para:

- Roteamento (Routing): Define URLs e como o servidor deve responder a elas (ex: quando alguém acessar /usuarios, execute tal função).
- Middleware: Permite interceptar requisições para fazer logs, autenticação, tratamento de erros ou compressão antes de chegar à lógica principal.
- Renderização de Respostas: Facilita muito o retorno de dados em formato JSON (muito usado em APIs), XML ou HTML.
- Validação de Dados: Ajuda a validar se os dados que o usuário enviou (no corpo da requisição) estão no formato correto.

### Pré requisito
- Instalar o gin web no projeto
```bash
go get github.com/gin-gonic/gin
```

### Exemplo
```go
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
```