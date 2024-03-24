package main

import (
	"fmt"

	"github.com/Guilherme129/Primeiro-API-Go/routes"
)

func main() {
	fmt.Println("Iniciando o servidor")
	routes.HandleRequest()
}
