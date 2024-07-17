package main

import (
	"fmt"       // Pacote para formatar e imprimir textos.
	"net/http"  // Pacote para fazer requisições HTTP.
	"os"        // Pacote para manipulação de entrada e saída do sistema operacional.
	"reflect"
)
	//No GoLang não temos a função "while". Utilizamos "for".
func main() {
	exibirNomes()
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
	var sites [4]string //Array com tamanho 4.
	//Os arrays tem tamanho fixo. Se declarou que tem tamanho 4 ele não pode ser modificado. Desto modo, utilizamos uma estrutura chamada de "slice". 
	sites[0] = "https://www.alura.com.br" 
	sites[1] = "https://www.google.com.br"
	sites[2] = "https://www.youtube.com"
	//sites[3] = Na posição 4.
	
	fmt.Println(sites)
	fmt.Println(reflect.TypeOf(sites))
	fmt.Println("O meu array tem o seguinte tamanho:", len(sites))
	fmt.Println("O meu arrya tem a capacidade  de:", cap(sites), "itens")

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

func exibirNomes() {
	nomes := []string{"Gabriel,", "Ana Laura,", "Izabel,", "Gesu,"} //Slice com tamanho 4.
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem o seguinte tamanho:", len(nomes), "itens") //len() função para saber o tamanho do slice.
	fmt.Println("O meu slice tem a capacidade  de:", cap(nomes), "itens") //cap() função para saber a capacidade do slice.

	nomes = append(nomes, "Feitosa") //A função append adiciona mais um no slice, fazendo com que seu tamanho fique 5.
	fmt.Println()

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("Agora meu slice tem o seguinte tamanho:", len(nomes), "itens")
	fmt.Println("O meu slice tem a capacidade  de:", cap(nomes), "itens")
	fmt.Println()
}

//Teste GIT.