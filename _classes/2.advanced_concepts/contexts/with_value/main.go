package main

import (
	"context"
	"fmt"
)

type ContextKey string

// O Context é criado como se fosse uma LinkedList, cada novo contexto instanciado referencia o contexto anterior:
// [background/main] <- [ctx1] <- [ctx2] <- [ctx3] <- ...
//
// Por isso, o algoritmo de busca de valores do Context é de O(n),
// então tome cuidado ao colocar os valores e se realmente precisam ser buscados na arvore de contexto
//
// Boas práticas do uso de valores dentro de contextos:
// 1. Se você gera dados na memória da sua goroutine, eles provavelmente não são bons candidatos
// 2. Dados devem ser imutáveis
// 3. Compartilhe dados simples ou simplificados
// 4. Os dados devem ser dados, não tipos com métodos
// 5. Utilize de dados que apenas decoram operações
//
// O que faz sentido passar para valores dentro dos contextos:
// 1. Traces ids
// 2. Start times
// 3. User ids / authorization / Request tokens
// 4. URLs
//
// Regras:
// 1. Não use strings para definir as chaves (key-value-pairs), pois podem ter colisões.
//    O valor não será sobrescrito porque os contextos são imutáveis, entretanto quando se usa a função `Value`,
//    o último valor encontrado é o que retorna:
//    - [background/main] <- [ctx1] <- [ctx2 value("anyKey", 10)] <- [ctx2 value("anyKey", 50)]
//    - Nesse caso, os valores 10 não foi sobrescrito pelo 50, entretanto o `Value("anyKey")` vai retornar o 50
//      por ser o primeiro caso encontrado
// 2. Use tipos privados customizáveis de primitivos para definir as chaves
//    - `type myCustomKey string`
//    - `context.WithValue(ctx, myCustomKey("anyKey"), value)`
// 3. Se precisar usar esses valores em outros pacotes que estão dentro do contexto em outros pacotes,
//    exporte funções que retornam diretamente o valor e não a chave
//    - `GetMyCustomKey(ctx) { return ctx.Value(myCustomKey("anyKey"))}`

type ctxKey string

func GetSchoolNameFromContext(ctx context.Context) string {
	return ctx.Value(ctxKey("schoolName")).(string)
}

func GetSchoolDescFromContext(ctx context.Context) string {
	return ctx.Value(ctxKey("schoolDesc")).(string)
}

// Caso queira passar o conteúdo da chave para o controle do solicitante,
// basta que a chave seja do mesmo tipo subjacente (underlying type)
func GetSchoolContext(ctx context.Context, key string) string {
	val := ctx.Value(ctxKey(key))
	if val != nil {
		return val.(string)
	}
	return ""
}

func doSomething(ctx context.Context, name, desc string) {
	ctx = context.WithValue(ctx, ctxKey("schoolName"), name)
	ctx = context.WithValue(ctx, ctxKey("schoolDesc"), desc)
	doSomethingElse(ctx)
}

func doSomethingElse(ctx context.Context) {
	fmt.Printf("School %s is %s:\n", GetSchoolNameFromContext(ctx), GetSchoolDescFromContext(ctx))
}

func main() {
	doSomething(context.Background(), "Harvard", "Muito legal")
}
