/*
	SQLite em Go precisa ter interoperação com C, ou seja, o Go não tem uma biblioteca nativa para SQLite.
	O Go tem na verdade uma stardard lib para interfaces de banco de dados, daí cada driver cria sua implementação.
	https://pkg.go.dev/database/sql

	O SQLite, como dito antes, precisa de uma interoperação com C e para isso usamos o pacote cgo que permite a
	interoperação entre Go e C. Com isso, é necessario ter o compilador de C instalado no local onde o binário será
	gerado, assim o Go consegue compilar o código C e gerar o binário com GO e C.
	Porem como a interoperação não é das melhores, precisamos consfigurar algumas coisas para que o Go consiga compilar.

	ATENÇÃO: A interoperação entre Go e C é uma feature avançada e deve ser usada com cuidado, pois pode causar problemas
	na performance(builds demorados e possivelmente quebrados, perda do cross compile, etc) e na segurança do código.

	Biblioteca e Artigos:
	-> "A sqlite3 driver that conforms to the built-in database/sql interface."
	-> https://github.com/mattn/go-sqlite3
	-> Artigo do Dave Cheney sobre a interoperação entre Go e C: https://dave.cheney.net/2016/01/18/cgo-is-not-go

*/

package main

import (
	"database/sql"
	"fmt"

	// Importando como blanked identifier para que o driver seja registrado e não seja necessário chamar diretamente
	_ "github.com/mattn/go-sqlite3"
)

/*
.

	sql.Open => Abre uma conexão com o banco de dados, o primeiro argumento é o driver que será usado e o segundo é a string de conexão
*/
func connectionToDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		panic(err)
	}
	return db
}

/*
.

	sql.Exec => Executa uma query que não retorna linhas, como INSERT, UPDATE, DELETE, CREATE, etc.

	O Resultado do método Exec é um objeto do tipo sql.Result que tem os seguintes métodos:
	- res.RowsAffected() => retorna o número de linhas afetadas pelo comando.
	- res.LastInsertId() => retorna o ID da última linha inserida.
*/
func createTableWithExec(db *sql.DB) {
	createTableSql := `
	CREATE TABLE IF NOT EXISTS foo (
		id integer not null primary key,
		name text
	);
	`
	res, err := db.Exec(createTableSql)

	if err != nil {
		panic(err)
	}

	fmt.Println(res.RowsAffected())
}

func insertTableValues(db *sql.DB) {
	insertTableSql := `
		INSERT INTO foo (name) VALUES ('Felipe');
	`

	res, err := db.Exec(insertTableSql)

	if err != nil {
		panic(err)
	}

	fmt.Println(res.RowsAffected())
}

/*
.

	Query() => Executa uma query que retorna linhas, como SELECT, SHOW, etc.
	QueryRow() => retorna uma única linha de uma consulta
	As queries podem receber argumentos que são passados como parâmetros para a query, através do `?`.

	Scan() => preenche as variáveis passadas por referencia como argumento com os valores retornados pela consulta,
	ou seja, ele faz o mapeamento dos valores retornados para as variáveis passadas por referencia.
*/
func queryAndTransformToStruct(db *sql.DB) {
	type user struct {
		ID   int64
		Name string
	}

	var u user
	querySql := `
		Select * from foo where id = ?;
	`
	if err := db.QueryRow(querySql, 1).Scan(&u.ID, &u.Name); err != nil {
		panic(err)
	}

	fmt.Println(u)
}

func main() {
	db := connectionToDb()

	createTableWithExec(db)
	insertTableValues(db)
	queryAndTransformToStruct(db)
}
