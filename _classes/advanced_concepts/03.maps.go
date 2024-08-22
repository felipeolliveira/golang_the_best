package advanced_concepts

import "fmt"

func Maps() {
	/*
		- Mapas são coleções de pares chave-valor, igual outras linguagens

		- Mapas podem ser inicializados de duas formas:
			myMap := map[string]string{} -> declaração literal
			myMap := make(map[string]string) -> make
	*/
	var invalidMap map[string]string
	fmt.Println(invalidMap == nil) // true
	// invalidMap["key"] = "value" // Isso vai dar erro, pois o mapa não foi inicializado

	validMap := make(map[string]string)
	fmt.Println(validMap == nil) // false

	validMap["Felipe"] = "Oliveira"
	validMap["Júlia"] = "Carminatti"
	fmt.Println(validMap) // map[Felipe:Oliveira Júlia:Carminatti]

	// Mapas podem ser inicializados com valores através de uma declaração literal
	mapWithInitialValues := map[string]string{
		"Felipe": "Oliveira",
		"Júlia":  "Carminatti",
	}
	fmt.Println(mapWithInitialValues) // map[Felipe:Oliveira Júlia:Carminatti]

	/*
		- Os valores de um mapa podem ser acessados pela chave -> map["key"]
			Quando um valor não existe, o valor retornado é o zero value do tipo do valor
			no caso de strings, o zero value é uma string vazia
			Por isso, o acesso aos valores de um mapa deve ser feito com um segundo retorno,
			que indica se a chave existe no mapa e por convenção é chamado de `ok`
	*/
	value, ok := validMap["Felipe"]
	fmt.Println(value, ok) // "Oliveira" true

	inexistentValue, ok := validMap["inexistentKey"]
	fmt.Println(inexistentValue, ok) // "" false

	/*
		- Para deletar um valor de um mapa, usamos a função delete(map, key)
	*/
	delete(validMap, "Felipe")
	fmt.Println(validMap) // map[Júlia:Carminatti]

	value, ok = validMap["Felipe"]
	fmt.Println(value, ok) // "" false

	/*
		- Para limpar o mapa sem alterar a capacidade do mapa, basta usar o clear(map)
			Vai manter o mapa, mas sem nenhum valor
			Esse método foi implementado para evitar a necessidade de recriar o mapa, já que podem
			ter casos que as chaves são NaN (Not a Number) e não podem ser comparadas
	*/

	clear(validMap)
	fmt.Println(validMap) // map[]

	validMap["fruit"] = "apple"
	validMap["color"] = "red"
	fmt.Println(validMap) // map[fruit:apple color:red]

	/*
	 - É possível iterar sobre um mapa usando o range
	 - A ordem dos elementos não é garantida
	 - É possível deletar elementos durante a iteração
	*/
	validMap["toRemove"] = "value"
	fmt.Println(validMap) // map[color:red fruit:apple toRemove:value]
	for key, value := range validMap {
		if key == "toRemove" {
			delete(validMap, key)
		}
		fmt.Println(key, value)
	}
	fmt.Println(validMap) // map[color:red fruit:apple]
}
