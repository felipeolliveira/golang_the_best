package main

import (
	"fmt"

	"github.com/Masterminds/squirrel"
)

/*
	Em alguns casos, é necessário construir queries SQL de forma dinâmica, ou seja, de acordo com a necessidade do usuário.
	Para isso, podemos utilizar um SQL Builder, que é uma ferramenta que nos ajuda a construir queries SQL de forma programática.

	Em Go, existem alguns pacotes:
		- squirrel => https://github.com/Masterminds/squirrel
		- goqu => https://github.com/doug-martin/goqu
		- go-sqlbuilder => https://github.com/huandu/go-sqlbuilder

	Todos esses pacotes são muito bons e possuem uma documentação bem completa, sendo o squirrel o mais antigo e já finalizado e o go-sqlbuilder o mais novo e com mais features.
*/

type Filters struct {
	Id       int64
	Name     string
	Email    string
	Username string
}

func build(f Filters) (string, []any) {
	/*
		Por padrão o squirrel faz o build da query com placeholder ? para os valores, mas podemos alterar isso para $1, :2, @3 etc, usando o método Placeholder.
	*/
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Select("*").From("users")
	or := squirrel.Or{}

	if f.Id > 0 {
		or = append(or, squirrel.Eq{"id": f.Id})
	}

	if f.Name != "" {
		or = append(or, squirrel.Like{"name": f.Name})
	}

	if f.Email != "" {
		or = append(or, squirrel.Eq{"email": f.Email})
	}
	if f.Username != "" {
		or = append(or, squirrel.Eq{"username": f.Username})
	}

	sql, args, err := builder.Where(or).ToSql()

	if err != nil {
		panic(err)
	}

	return sql, args
}

func main() {
	filters := Filters{
		Id:       1,
		Name:     "John",
		Email:    "john@email.com",
		Username: "john",
	}
	sql, args := build(filters)

	fmt.Println(sql)
	fmt.Println(args)
}
