package advanced_concepts

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

/*
  - Generics são um recurso que permite a criação de funções, estruturas e interfaces que aceitam tipos genéricos, bem semelhante a outras linguagens como Java, C# e TypeScript
    São usados em funções ou structs
    Generics precisam sempre de um constraint junto da definicao do tipo genérico
    foo[<Nome_Generic> <Constraints>] (arg <Nome_Generic>) {}
    foo[T any] (arg T) {}

  - Generics `any` tem funcionamento diferente de `any` como tipo.
    `any` como tipo: é realizada a checagem de tipos em tempo de execução
    `any` como generics: é realizada a checagem de tipos em tempo de compilação. O compilador gera uma versão da função para cada tipo que é chamada
*/
func foo[T any](arg T) {
	fmt.Println("foo generic", arg)
}

/*
  - Há vários tipos de constraints que podem ser usados
    any: aceita qualquer tipo
    int, string, float64, etc: aceita apenas o tipo específico
    interface: aceita qualquer tipo que implemente a interface
    comparable: aceita qualquer tipo que possa ser comparado, ou seja, que implemente os métodos == e !=
*/
type MyConstraint interface {
	Foo()
}

func fooWithMyConstraint[T MyConstraint](arg T) {
	fmt.Println("foo with constraint", arg)
}

// Agora a função fooWithMyConstraint aceita o MyType porque ele implementa a interface MyConstraint
type MyType struct{}

func (MyType) Foo() {}

/*
  - Dentro dos generics, é possível utilizar multiplos constraints usando uma interface que possua os contraints desejados, nesse caso, os contraints ou tipos são separados por pipe (|) semelhante ao Typescript. Há duas formas de definir esses constraints multiplos:
    MyFunc[T interface { Foo | Bar }](): aceita a criação da interface inline
    MyFunc[T MyConstraint](): aceita uma interface já criada com as opções dentro
    -> MyConstraint interface { Foo | Bar }
*/
type MultipleConstraits interface {
	int | string
}

func multipleConstraints[T MultipleConstraits](arg T) {}

/*
  - Quando se tem multiplos constraints, o funcionamento é estrito, ou seja, o tipo passado como argumento deve ser um dos tipos definidos no constraint.
    Caso queira aceitar um tipo customizado mas que o `core type` dele esteja dentro das opções dos generics, é necessário usar o `~` antes do tipo informado

  - inclusive, o Go tem um pacote com essas constraints que são usadas internamente, com o objetivo de facilitar a criação de generics com `core types`
    `golang.org/x/exp/constraints`
*/
type MyTypeString string

var myStr MyTypeString = "string"

type MultipleConstraitsWithCoreType interface {
	// Se não for passado o `~` antes do tipo customizado, o compilador irá acusar erro
	// Pois nesse caso, o tipo `MyTypeStrin` não é um dos tipos definidos no constraint, mas internamento, lá no `core type` ele é uma string que é um dos tipos definidos no constraint
	int | ~string | constraints.Float
}

func multipleConstraintsWithCoreType[T MultipleConstraitsWithCoreType](arg T) {}

/*
	- Structs podem receber generics para atribuir as suas propriedades

	- Métodos de Structs também podem receber genérics

	- LIMITAÇÃO: Não é possível usar generics em métodos de structs:
		Uma maneira de tentar driblar isso é usar os argumentos com os mesmos tipos do receiver
		func (myStruct[T]) makeSomething[A string](a A) {} -> erro de compilação
*/

type myStructWithGeneric[T string] struct {
	Foo T
}

func (ms myStructWithGeneric[T]) makeFooWithArgs(a T) {
	fmt.Println(ms.Foo, a)
}
func (ms myStructWithGeneric[T]) makeFoo() {
	fmt.Println(ms.Foo)
}

/*
	- LIMITAÇÃO: Não é possível recuperar as informações de propriedades e métodos em tempo de execução mesmo que as interfaces sejam implementadas corretamente
*/

type Someone struct {
	Name string
}

func (Someone) Run() string {
	return "Running"
}

type Car struct {
	Name string
}

func (Car) Run() string {
	return "Running"
}

func makeRun[T interface{ Someone | Car }](t T) {
	// Mesmo que a interface seja implementada corretamente, ainda não é possível acessar as informações de ambas as structs
	// fmt.Println(t.Name)  // erro de compilação
	// fmt.Println(t.Run()) // erro de compilação
}

// Uma maneira de driblar isso e pelo menos conseguir acessar os métodos, é passar uma constraint já construida com o método em comum nas structs
type SomeoneOrCarConstraint interface {
	Someone | Car
	Run() string
}

func makeRunWithMethodInsideConstraint[T SomeoneOrCarConstraint](t T) {
	// Agora é possível acessar o método comum nas structs
	// fmt.Println(T.Name)  // erro de compilação
	fmt.Println(t.Run())
}

func Generics() {
	foo("")
	foo(1)
	foo([]int{1, 2, 3})

	fooWithMyConstraint(MyType{})

	multipleConstraints(1)
	multipleConstraints("string")

	multipleConstraintsWithCoreType(myStr)

	makeRun(Someone{Name: "Someone"})

	ms := myStructWithGeneric[string]{Foo: "foo"}
	ms.makeFooWithArgs("bar")
	ms.makeFoo()

	makeRun(Car{Name: "Car"})
	makeRunWithMethodInsideConstraint(Car{Name: "Car"})
}
