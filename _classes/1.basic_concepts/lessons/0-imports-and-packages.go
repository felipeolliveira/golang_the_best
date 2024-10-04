/*
  - Pacotes devem ser a primeira declaração de um arquivo
    Apenas comentários podem ser declarados antes de um pacote

  - O Pacote é divido pelo nome do diretório onde o arquivo está
    Geralmente o nome do pacote é o mesmo nome do diretório

  - Pacote main não pode ser importado
    É o pacote principal de um programa Go, onde o programa começa a execução

  - Pacotes dentro da pasta `internal` só podem ser importados por pacotes dentro da mesma árvore de diretórios (irmãos)
    É uma forma de garantir que pacotes internos não sejam importados por pacotes externos
    ├-pacote
    ├		└- internal                 // pacote interno
    ├		├			└- funcao-interna.go
    ├		├- a.go                     // pode usar o pacote interno
    ├ 	├- b.go                     // pode usar o pacote interno
    └- main.go                      // não pode usar o pacote interno
*/
package lessons

/*
	- Importação padrão de pacotes, onde o nome do pacote já é a importação: "basic_concepts"

	- Import Statement: import ("go_training/basic_concepts")
		Geralmente é usado para importação de multiplos pacotes

	- Import Alias: import basic_concepts "go_training/basic_concepts"
		É usado para importar um pacote com um nome diferente, geralmente usado para evitar conflitos de nomes

	- Import Blank Identifier: import _ "go_training/basic_concepts"
		É usado para importar um pacote sem usar o pacote, apenas para executar o init() do pacote

	- Import Dot: import . "go_training/basic_concepts"
		É usado para importar um pacote e usar as funções do pacote sem usar o nome do pacote
		Não é recomendado, pois pode causar confusão no código
*/
import "fmt"

func Imports() {
	fmt.Println("go" + "lang")
}
