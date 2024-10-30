package main

import "fmt"

func main() {
	felipe := Person{"Felipe"}

	felipe.talk()
	// Essa seria a forma "correta" para chamar um metodo de uma struct implementada por interface
	// já que o receiver recebeu um ponteiro para a struct, porem o Go faz uma excessão e consegue usar apenas o `.talk()` diretamente
	(&felipe).talk()

	saySomething(&felipe)
	// saySomething(felipe) -> Nesse caso, não tem como, precisa passar o ponteiro para que a funcão aceite a interface `Humans`
}

type Humans interface {
	talk()
}

type Person struct {
	name string
}

func (p *Person) talk() {
	fmt.Println("Olá, eu sou pessoa chamada:", p.name)
}

func saySomething(h Humans) {
	h.talk()
}
