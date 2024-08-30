package main

/*
	Existem alguns pacotes para drivers do Postgres:
	- github.com/lib/pq => de acordo com a starndard lib
	- github.com/jackc/pgx => algumas funcionalidades a mais que o lib/pq.
		- as funcionalidades a mais estão fora da starndard lib, então é uma implementação própria.
		- integração com outra funcionalidade do sqlc, que é um gerador de código SQL.

	Ambos são bem avaliados, mas nesse caso, irei utilizar o jackc/pgx.
*/

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func connectToBD(ctx context.Context) *pgxpool.Pool {
	pgUrl := "postgres://user:password@localhost:5450/db_main"

	/*
		Diferente do MySQL, o driver do Postgres não cria um pool de conexões automaticamente, precisa usar um método específico para isso.
		Basta usar o pacote pgxpool.

		- pgx.Connect => cria uma conexão com o banco de dados.
		- pgxpool.New => cria um pool de conexões com o banco de dados.
	*/
	conn, err := pgxpool.New(ctx, pgUrl)

	if err != nil {
		slog.Error("postgres", "connectToDB", "Unable to connect to the database", "err", err.Error())
		os.Exit(1)
	}

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	slog.Info("postgres", "connectToDB", "Successfully connected to the database")
	return conn
}

func createTables(ctx context.Context, conn *pgxpool.Pool) {
	query := "create table if not exists products (id bigserial primary key, name varchar(255), price bigint);"

	/*
		O Exec do pgx necessida do contexto, que é uma forma de cancelar a execução de uma query, caso seja necessário.
		Diferente da standard lib, que não obriga o uso do contexto no Exec(), apenas no ExecContext().
	*/
	if _, err := conn.Exec(ctx, query); err != nil {
		slog.Error("postgres", "createProducts", err.Error())
	} else {
		slog.Info("postgres", "createProducts", "Table products created successfully")
	}
}

func insertProduct(ctx context.Context, conn *pgxpool.Pool, name string, price int) {
	/*
		No pgx, o placeholder é $1, $2, $3, etc, diferente do ? do go-sql-driver/mysql.
		Assim, pode ser reaproveitado os valores atribuidos nos placeholders.
	*/
	query := "insert into products (name, price) values ($1, $2);"

	if _, err := conn.Exec(ctx, query, name, price); err != nil {
		slog.Error("postgres", "error in insertProduct", err.Error())
	} else {
		slog.Info("postgres", "insertProduct", fmt.Sprintf("Product %s with price %d inserted successfully", name, price))
	}
}

type product struct {
	id    int
	name  string
	price int
}

func selectAllProducts(ctx context.Context, conn *pgxpool.Pool) []product {
	res := []product{}

	query := "select * from products;"
	/*
		O método Query retorna um objeto do tipo Rows, que é um conjunto de linhas retornadas por uma consulta.
	*/
	rows, err := conn.Query(ctx, query)
	if err != nil {
		slog.Error("postgres", "error in selectAllUsers", err.Error())
	}

	/*
		Não pode esquecer de fechar o rows, caso contrário, a conexão com o banco de dados não será fechada
	*/
	defer rows.Close()

	/*
		Para percorrer as linhas retornadas, é necessário usar o método Next, que retorna true se houver uma linha a ser lida e false se não houver.
		Dentro do loop, é possível usar o método Scan para ler os valores de cada coluna da linha atual e montar a estrutura de dados desejada,
		sempre passando o endereço da variável que deseja preencher.
	*/
	for rows.Next() {
		var p product
		if err := rows.Scan(&p.id, &p.name, &p.price); err != nil {
			slog.Error("postgres", "error in selectAllProducts", err.Error())
		}
		res = append(res, p)
	}

	return res
}

func main() {
	/*
		o pgx/v5 por padrão já usa obriga o contexto
	*/
	ctx := context.Background()
	conn := connectToBD(ctx)

	createTables(ctx, conn)
	insertProduct(ctx, conn, "Apple", 50)

	users := selectAllProducts(ctx, conn)
	slog.Info("postgres", "products", fmt.Sprintf("%+v", users))

	defer conn.Close()
}
