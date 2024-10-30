package lessons

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/*
	JSON é um pacote do encoding que permite a codificação e decodificação de dados no formato JSON.
	As principais funções são:
	- Marshal: converte um tipo para um json em bytes (string)
	- Unmarshal: converte os bytes de um json para um tipo passado por ponteiro
	- NewEncoder: cria um encoder para escrever um json em um writer
	- NewDecoder: cria um decoder para ler um json de um reader

	JSON tem struct tags que são usadas para mapear os campos da struct para o JSON.
	- `json:"name"`: mapeia o campo da struct para o campo do JSON
	- `json:"-"`: ignora o campo da struct no JSON
	- `json:"name,omitempty"`: mapeia o campo da struct para o campo do JSON e ignora se o campo for vazio
	- `json:",omitempty"`: ignora o campo do JSON se o campo da struct for vazio
	- `json:"name,string"`: mapeia o campo da struct para o campo do JSON e converte o tipo para string
	- `json:"name,omitempty,string"`: mapeia o campo da struct para o campo do JSON e converte o tipo para string e ignora se o campo for vazio
*/

type Fruit struct {
	Name   string `json:"name"`
	Color  string `json:"color"`
	Rate   int    `json:"rate"`
	IsGood bool   `json:"isGood"`
}

func JsonParsing() {
	marshal()
	unmarshal()
	encoder()
	decoder()
}

/*
Marshal faz o parsing do tipo para um json em bytes (string)
{"key": "value",}
*/
func marshal() {
	// Marshal
	apple := Fruit{"Apple", "Red", 5, true}
	appleJson, err := json.Marshal(apple)

	if err != nil {
		panic(err)
	}

	fmt.Println(apple)
	fmt.Println(string(appleJson))
}

/*
Unmarshal faz o parsing dos bytes de um json para um tipo passado por ponteiro
{"key": "value",}
*/
func unmarshal() {
	// Marshal
	banana := Fruit{}
	bananaBytes := []byte(`{"Name":"Banana","Color":"Yellow","Rate":0,"IsGood":false}`)
	err := json.Unmarshal(bananaBytes, &banana)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bananaBytes))
	fmt.Println(banana)
}

func encoder() {
	orange := Fruit{"Orange", "orange", 10, true}

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(orange)
}

func decoder() {
	melonBytes := []byte(`{"Name":"Melon","Color":"Green","Rate":0,"IsGood":true}`)

	dec := json.NewDecoder(strings.NewReader(string(melonBytes)))

	for dec.More() {
		token, err := dec.Token()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v ", token)
	}

	token, err := dec.Token()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", token)
}
