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
	if(comando == 1){ //No GoLang não necessitamos colocar parênteses no IF. Temos que sempre colocar uma booleana no IF. Algo que retorne true ou false.
		fmt.Println("Monitorando...")
	}else if(comando ==2){
		fmt.Println("Exibindo logs...")
	}else if(comando ==0){
		fmt.Println("Saindo do programa...")
	}else {
		fmt.Println("Não conheço este comando!")
	}
	fmt.Println()
}
