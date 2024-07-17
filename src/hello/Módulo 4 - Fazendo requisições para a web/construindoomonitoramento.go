package main

import (
	"fmt"
	"os"
)

func main() {
	exibirDadosIniciais() //Chamar a função escrita embaixo, func exibirDadosiniciais()
	exibirMenu() //Chamar a função func exibirMenu()
	//entradaDados() //Chamar a função func entradaDados()

	//Declarando a variável int utilizada na função entradaDados()
	//Primeira maneira (completa):
	/*var comando int
	comando = entradaDados()*/

	//Segunda maneira (simplificada):
	comando := entradaDados() //Equivale a exibir func entradaDados()
	
	//Utilizando a função switch (trocar).
	switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")	
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando!")
			os.Exit(-1)
	}
	//Em GoLang não necessitamos colocar o "break".
}

func exibirDadosIniciais() { //Em GoLang é comum utilizar a primeira letra em minúsculo e a outras em maiúsculo.
	//Dados iniciais.
	nome := "Gabriel"
	versao := 1.1
	fmt.Println("Olá, senhor "+ nome +"!")
	fmt.Println("O programa está na versão", versao)
	fmt.Println()	
}

func exibirMenu() {
	//Menu de opções.
	fmt.Println("1 - Monitoramento de site")
	fmt.Println("2 - Exibir os logs")
	fmt.Println("0 - Sair do programa")
	fmt.Println()
}

func entradaDados() int {
	//Entrada de dados.
	fmt.Println("Digite a opção desejada: ")
	var comandoUsuario int
	fmt.Scan(&comandoUsuario)
	fmt.Println("O comando escolhido foi:", comandoUsuario)
	fmt.Println()

	return comandoUsuario
}
	//Saída do programa.
/*func saidaPrograma() {
	fmt.Println("Saindo do programa.")
	//Agora, a função saidaPrograma será chamada apenas quando o usuário escolher a opção 0, e a variável comandoSaida foi removida.
	//Eliminando o erro de variável não utilizada.
}*/
func iniciarMonitoramento() { //Será necessário fazer uma requisão web para testar o site. No GoLang temos o pacote: 
	fmt.Println("Monitorando...")
	// site := "https//www.alura.com.br"
	// resp, err := http.Get(site)
	fmt.Println()
}

