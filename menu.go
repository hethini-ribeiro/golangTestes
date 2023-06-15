package main

import (
	"fmt"
)

func main() {
	nome := "hethini"
	var versao = 1.1
	fmt.Println("Olá, sta ", nome)
	fmt.Println("Este programa na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
}
