package main

import (
	"fmt"	
)

func main() {
	//Dados iniciais.
	nome := "Gabriel"
	versao := 1.1
	fmt.Println("Olá, senhor "+ nome +"!")
	fmt.Println("O programa está na versão", versao)
	fmt.Println()

	//Menu de opções.
	fmt.Println("1 - Monitoramento de site")
	fmt.Println("2 - Exibir os logs")
	fmt.Println("0 - Sair do programa")
	fmt.Println()

	//Entrada de dados.
	fmt.Println("Digite a opção desejada: ")
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi:", comando)
	fmt.Println()

	//Utilizando a função If (se).
	switch comando {
		case 1:
			fmt.Println("Monitorando...")
		case 2:
			fmt.Println("Exibindo logs...")	
		case 0:
			fmt.Println("Saindo do programa...")
		default:
			fmt.Println("Não conheço este comando!")
	}
	//Em GoLang não necessitamos colocar o "break".
}

