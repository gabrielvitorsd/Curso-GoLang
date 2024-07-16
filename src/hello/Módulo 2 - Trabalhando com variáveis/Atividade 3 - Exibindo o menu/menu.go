package main

import (
	"fmt"
	
)

func main() {
	//Dados inicias
	nome := "Gabriel"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão:", versao)
	fmt.Println()

	//Menu de opções
	fmt.Println("1 - Monitoramento do site")
	fmt.Println("2 - Exibir os logs")
	fmt.Println("0 - Sair do programa")
	fmt.Println()

	//Entrada de dados pelo usuário = input
	fmt.Println("Digite a opção desejada:")
	var comando int //O comando que o usuário digita sendo 1,2 e 0 é um inteiro.
	fmt.Scanf("%d", &comando) //O especificador "%d" é usado para ler um valor inteiro decimal e o & é usado para obter o endereço de uma variável.
	fmt.Println("O endereço da minha variável comando é", &comando)
	fmt.Println("O comando escolhido foi:", comando)
}



