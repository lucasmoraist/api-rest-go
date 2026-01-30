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