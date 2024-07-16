package main //Programas em GO, podem ter 1 ou mais pacotes. Assim, declarando o pacote 'main'. 
//Ao especificar 'package main', você está dizendo ao compilador que este é o ponto de entrada do programa e que o arquivo será usado para gerar um executável.
import "fmt"

func main() { //Todo pacote principal, tem a função principal 'main'. 
//Para imprimir uma mensagem nós utilizamos a função Println. 
//Porém, é necessário importar um pacote específico para usá-la que é o "fmt". fmt = format. importação = import "fmt" 
	fmt.Println("Hello World! Teste do - go run hello.go -") 
}
