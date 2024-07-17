package main

import (
	"fmt"
	"net/http"
	"os"
)

//1. Exibir a entrada de dados.
//2. Menu de opções.
//3. Fazer a entrada de dados.
//4. Iniciar o monitoramento.

func main() {
	exibirDadosIniciais()
	for {
		exibirMenu()
	
		comando := entradaDados()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 3:
			fmt.Println("Sair do programa")
			os.Exit(0)
		default:
			fmt.Printf("Comando não reconhecido")
		}
	}
}

func exibirDadosIniciais() {
	fmt.Println()
	nome := "Ana Laura"
	versao := 1.1
	fmt.Println("Olá, sra.", nome)
	fmt.Println("O programa está na versão", versao)
	fmt.Println()
}

func exibirMenu() {
	fmt.Println("1. Monitorando...")
	fmt.Println("2. Exibir logs...")
	fmt.Println("3. Sair do programa...")
}

func entradaDados() int {
	fmt.Println()
	fmt.Println("Digite a opção desejada: ")
	var comandoUsuario int
	fmt.Scan(&comandoUsuario)
	fmt.Println("O comando escolhido foi:", comandoUsuario)
	fmt.Println()

	return comandoUsuario
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://www.youtube.com/"
	resp, _ := http.Get(site)
	fmt.Println()
	fmt.Println(resp)
	fmt.Println()

}