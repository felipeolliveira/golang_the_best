package main

/*
Evitando o CGO para SQLite...

	Há pacotes que permitem a comunicação com o SQLite utilizando uma implementação pura em Go.
	Porém há duas abordagens para se comunicar com o SQLite sem utilizar o CGO:

	1. Utilizando as interfaces da stardard library `database/sql`
		- modernc.org/sqlite
		- gitlab: https://gitlab.com/cznic/sqlite

	2. Não utilizando as interfaces da stardard library `database/sql`
		- zombiezen.com/go/sqlite: é um fork crawshaw.io/sqlite Não usa as interfaces da starndard library
		- github: github.com/zombiezen/go-sqlite

	Porque usar a opção 2?
	Há um bom artigo sobre o porque não usar as interfaces padrões do `database/sql` para o SQLite:
		- O SQLite é um banco de dados diferente dos outros bancos de dados, pois é um banco de dados de arquivo.
		- Não faz sentido ser orientado a conexões, pois o SQLite é um banco de dados que roda no mesmo processo que a aplicação.
		- O SQLite suporta `parameter names`, que ajuda na identificação dos parâmetros de uma query, sem precisar usar `?`.
			Se seguisse a interface padrão, não seria possível usar essa feature do SQLite.
		- O SQLite suporta `nested transactions` e `savepoints`, que não são suportados pela interface padrão.
		- Mais detalhes em https://crawshaw.io/blog/go-and-sqlite
*/

import (
	// Também é necessário importar como blanked identifier para que o driver seja registrado
	"fmt"

	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

/*
.

	A conexão com o banco de dados é feita de forma semelhante, mas não identica, pois não é mais necessário registrar o driver.
	Agora o retorno da função é um `*sqlite.Conn` e não um `*sql.DB`.

	path: ":memory:"
	É possível criar um banco de dados em memória, passando `:memory:` como argumento para a função `sqlite.OpenConn`.
	Isso é útil para testes, pois o banco de dados é criado em memória e não é persistido no disco.
*/
func connectionToDB() *sqlite.Conn {
	conn, err := sqlite.OpenConn("./sqlite-pure-go.db", sqlite.OpenReadWrite|sqlite.OpenCreate)
	if err != nil {
		panic(err)
	}

	return conn
}

func main() {
	conn := connectionToDB()
	defer conn.Close()

	// Execute a query.
	err := sqlitex.ExecuteTransient(conn, "SELECT 'hello, world';", &sqlitex.ExecOptions{
		ResultFunc: func(stmt *sqlite.Stmt) error {
			fmt.Println(stmt.ColumnText(0))
			return nil
		},
	})
	if err != nil {
		panic(err)
	}
}
