package main

import (
	"fmt"       // Pacote para formatar e imprimir textos.
	"net/http"  // Pacote para fazer requisições HTTP.
	"os"        // Pacote para manipulação de entrada e saída do sistema operacional.
	
)
	//No GoLang não temos a função "while". Utilizamos "for".
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
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando!")
			os.Exit(-1)
		}	
	}
	 //Chamar a função escrita embaixo, func exibirDadosiniciais()
	 //Chamar a função func exibirMenu()
	//entradaDados() //Chamar a função func entradaDados()

	//Declarando a variável int utilizada na função entradaDados()
	//Primeira maneira (completa):
	/*var comando int
	comando = entradaDados()*/

	//Segunda maneira (simplificada):
	 //Equivale a exibir func entradaDados()
		
	//Utilizando a função switch (trocar).

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
	site := "https://www.alura.com.br" 
	// resp, err := http.Get(site) //Faz uma requisição para o site. A função get tem um retorno que é a resposta (resp) do site.
	resp, _ := http.Get(site) //O underscore (_) é usado em Go para ignorar o segundo valor retornado pela função http.Get(site).
	//A função http.Get(site) retorna 2 valores.
	fmt.Println()
	fmt.Println(resp)
	fmt.Println()

	if resp.StatusCode == 200 { //O 200 vem da resposta o site está funcionando quando dá o valor de 200. Exemplo: &{200 OK 200 HTTP/2.0 2 0
		fmt.Println("O site", site, "foi carregado com sucesso.") //Status == 200.
	}else {
		fmt.Println("O site", site, "não está funcionando. Status code:", resp.StatusCode) //Status != 200
	}
	fmt.Println()
}

//Teste GIT.