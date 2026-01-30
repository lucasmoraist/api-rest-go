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