package advanced_concepts

import "fmt"

/*
	# Interfaces

	- Interfaces é uma forma de definir um contrato para um tipo, semelhante a abstract classes em outras linguagens.
	- Uma interface é um tipo e pode ser criada com a palavra-chave `type` e `interface`.
	- `any` em Go é representado por `interface{}`, ou seja, uma interface vazia.

	- Por convenção, o nome de uma interface é o nome do método seguido de `er`, por exemplo, `Reader`, `Writer`, `Closer`, `Stringer`, etc.
		Na maioria dos casos, respondendo ao serviço que aquela interface faz: `Aquele que gera (Generator)`, `Aquele que lê (Reader)`, `Aquele que escreve (Writer)`, `Aquele que fecha (Closer)`, `Aquele que converte para string (Stringer)`, etc.
`*/

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

/*
  - IMPORTANTE:
  - Para Go, uma struct implementa uma interface se ela tiver todos os métodos da interface, independente se a 		interface é conhecida por quem implementa.
    Não existe palavra-chave `implements` em Go, então a implementação é sempre implícita.
*/
func (Dog) Speak() string {
	return "Woof!"
}

func whatDoesTheAnimalSay(a Animal) {
	fmt.Println(a.Speak())
}

/*
  - Caso uma variavel instancie uma struct com metodo com receiver por ponteiro que implementa uma interface, e chame um método da interface, o método da struct será chamado mesmo a variavel sendo nil.
    Isso acontece porque o Go chama o método diretamente da struct, e não da variavel.
    Só irá acontecer um panic se for chamado uma propriedade da struct que não foi inicializada.
*/
type Cat struct {
	Name string
}

func (c *Cat) Speak() string {
	// c.Name = "Tom" // panic: runtime error: invalid memory address or nil pointer dereference
	// Como não está sendo chamado nenhuma propriedade da struct, não irá acontecer um panic.
	return "Meow!"
}

func Interfaces() {
	dog := Dog{Name: "Rex"}

	fmt.Println(dog.Speak())  // Woof!
	whatDoesTheAnimalSay(dog) // Woof!
	// Mesmo que a função `whatDoesTheAnimalSay` não saiba que `dog` é um `Dog`, ela sabe que `dog` implementa a interface `Animal` de forma implicita, pois o `Dog` implementa todas as funcões dessa interface `Animal`, logo `Dog` é um `Animal`.

	var a Animal           // nil
	var cat *Cat           // nil
	a = cat                // Cat -> Animal -> nil
	fmt.Println(a)         // nil
	fmt.Println(a.Speak()) // meow!
	// Não irá acontecer um panic, pois o Go chama o método diretamente da struct por ponteiro, e não da variavel.
	// A não ser que dentro do método da struct seja chamado uma propriedade da struct que não foi inicializada.
}
