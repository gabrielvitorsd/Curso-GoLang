package main 

import "fmt"

func main() {
    //1. Declarando a variável nome.
    var nome1 string = "Gabriel1" //Declaração de variável completa. String sempre está entre aspas.
    var nome2 = "Gabriel2" //Declaração de variável resumida. String sempre está entre aspas.
    nome3 := "Gabriel3" //Declaração de variável mais resumida. String sempre está entre aspas.
    
    //2. Declarando a variável versão.
    var versao1 float32 = 1.1 
    var versao2 = 1.2 
    versao3 := "15.07.24" //Leu como uma string visto que Você deve usar apenas um ponto decimal para separar a parte inteira da parte fracionária. 
    //Se você precisa armazenar uma versão como 15.07.24, deve armazená-la como uma string. Ou pode usar 15.0724
    
    //3. Declarando a variável idade.
    var idade1 int = 27
    var idade2 int //Quando não recebe valor o Go entende a variável como "zero".
    idade3 := 27
    //1.1
    fmt.Println("Hello World,", nome1,"!")
    fmt.Println("Hello World, " + nome2 + "!")
    fmt.Println(fmt.Sprintf("Hello World, %s!", nome3))
    //2.1
    fmt.Println("Este programa está na versão ", versao1)
    fmt.Println(fmt.Sprintf("Este programa está na versão ", versao2))
    fmt.Println("Este programa está na versão", versao3)
    //3.1
    fmt.Println("Olá, sr.", nome1, "sua idade é", idade1)
    fmt.Println("Olá, sr.", nome1, "sua idade é", idade2)
    fmt.Println("Olá, sr.", nome1, "sua idade é", idade3)
}
 

