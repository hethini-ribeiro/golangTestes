package main

import (
	"fmt"
	"os"
)

func main() {
	exibeIntroducao()
	exibeMenu()

	comando := leParametro()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Encerrando programa ...")
		os.Exit(0)
	default:
		fmt.Println("Comando não identificado!")
		os.Exit(-1)
	}

}

func exibeIntroducao() {
	nome := "hethini"
	var versao = 1.1
	fmt.Println("Olá, sta ", nome)
	fmt.Println("Este programa na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func leParametro() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
	return comando
}
