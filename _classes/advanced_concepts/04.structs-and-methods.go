package advanced_concepts

import (
	"encoding/json"
	"fmt"
)

/*
	- Tipos de dados compostos sempre usam a palavra-chave `type` para serem criados.
		Qualquer tipo de dado pode ser um tipo de dado composto.
*/

type MyString string

/*
  - Structs são tipos de dados que podem ser usados para agrupar diferentes tipos de dados.
    Uma struct é um tipo e pode ser criada com o `type` e `struct` keywords.
    Structs não são classes, mas são bem parecidas por terem propriedades e métodos.
*/
type Person struct {
	ID   uint64
	Name string
	Age  uint8
}

/*
  - Para criar métodos é necessário usar um `receiver` que é um parâmetro especial que é passado para o método.
    É declarada como uma função normal, mas com um parâmetro antes do nome da função.
    Exemplo: func (<Struct>) <func_name>() {}

  - O receiver pode ser um ponteiro ou um valor. Com ele é possível acessar os campos da struct.
    Quando a struct é passada como valor, uma cópia é feita e qualquer alteração não será refletida no original.
    Quando a struct é passada como ponteiro, a referência é passada e qualquer alteração será refletida no original.
    Quando não é necessário usar o valor da struct, pode omitir a variavel do receiver, passando apenas o tipo.

    Por convenção, o nome do receiver é a primeira letra do tipo da struct em minúsculo.
    Por convenção, quando um método altera o estado da struct, o receiver deve ser um ponteiro. Por consequência,
    todos os métodos da struct devem ser ponteiros.

    Não é possível adicionar metodos de structs fora do pacote que a struct foi definida.
*/
func (MyString) SayHello() {
	fmt.Println("Hello!")
}

// var str MyString = "any string"
// str.SayHello() -> Hello!

func (p Person) GetIDWithPrefix() string {
	return fmt.Sprintf("ID: %d", p.ID)
}

// A partir desse momento, eu precisaria alterar o metodo (p Person) para (p *Person)
// em todos os métodos da struct Person, inclusive os que não alteram o estado da struct.
func (p *Person) SetName(name string) {
	p.Name = name
}

/*
- Embbeded structs são structs que são usadas como campos de outras structs, bem semelhante a herança, mas os campos são de fato copiados para a struct que os contém.
*/
type Developer struct {
	Person
	Language string
}

/*
  - Struct Tags são strings que podem ser adicionadas aos campos de uma struct para fornecer metadados sobre o campo.
  - Struct Tags são usadas para serializar e deserializar structs. Qualquer pacote que aceite structs serializadas e deserializadas pode usar essas tags. Inclusive é possível criar pacotes personalizados que aceitam essas tags.
    Exemplo: `json:"name"`, `xml:"name"`, `yaml:"name"`, `bson:"name"`, `db:"name"`, `validate:"name"`, etc.
*/
type UserWithTags struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func StructsAndMethods() {
	/*
		- Variaveis que usam Structs podem ser declarados e inicializados usando a sintaxe de struct literal, omitindo o nome dos campos, porém, a ordem dos valores deve ser a mesma dos campos.
		- Caso o valor de um campo não seja fornecido, ele será inicializado com o valor zero do tipo do campo.
	*/
	john := Person{
		ID:   1,
		Name: "John Doe",
		Age:  30,
	}

	john2 := Person{2, "John Doe", 30}
	john3 := Person{ID: 3, Name: "John Doe"} // Age = 0
	julia := Person{ID: 4, Name: "Júlia Carminatti Ferrari", Age: 29}
	johnDev := Developer{
		Person:   Person{ID: 5, Name: "John Doe", Age: 30},
		Language: "Go",
	}
	johnWithTags := UserWithTags{ID: 1, Name: "John Doe", Age: 30}

	john.GetIDWithPrefix()     // ID: 1
	john.SetName("John Smith") // Name: John Smith

	fmt.Println(john, john2, john3, johnDev)
	fmt.Println(julia)

	johnWithTagsJson, _ := json.Marshal(johnWithTags)
	johnJson, _ := json.Marshal(john)
	fmt.Println(string(johnWithTagsJson)) // {"id":1,"name":"John Doe","age":30}
	fmt.Println(string(johnJson))         // {"ID":1,"Name":"John Doe","Age":30}
}
