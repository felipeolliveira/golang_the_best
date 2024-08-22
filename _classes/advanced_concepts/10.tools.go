package advanced_concepts

import "fmt"

/*
	- Tools são programas que ajudam a automatizar tarefas comuns de desenvolvimento
	- O Go possui várias ferramentas que ajudam a desenvolver, testar e manter código
		gopls: Language Server Protocol para Go
		gofmt: Formata o código
		...
		golang.org/x/tools
*/

type Example struct {
	Name string
	Temp float64
	List []string
}

func GoTools() {
	example := Example{
		Name: "",
		Temp: 0,
		List: []string{},
	}
	fmt.Println(example)
}
