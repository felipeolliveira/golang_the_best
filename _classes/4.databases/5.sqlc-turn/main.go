package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/felipeolliveira/golang_the_best/_classes/databases/sqlc-turn/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

/*
	Irei utilizar dois pacotes para lidar automações no processo de usar banco de dados SQL:
	- Tern => https://github.com/jackc/tern => utilizado para gerar migrations do banco de dados.
		- alternativas
		- goose => https://github.com/pressly/goose
		- golang-migrate => https://pkg.go.dev/github.com/golang-migrate/migrate/v4
	- SQLc => https://sqlc.dev/ => utilizado para gerar as queries, models e abstrações para conexão com o banco de dados.

	Em outras palavras:
	- O tern é para migrations!
	- O sqlc é quase um ORM
	- Juntando as duas ferramentas, chegamos perto de um ORM, mas sem as desvantagens de um ORM.

*/

func connectToBD(ctx context.Context) *pgxpool.Pool {
	pgUrl := "postgres://user:password@localhost:5450/sqlc_turn"
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

func listAllAuthors(ctx context.Context, queries *db.Queries) []db.Author {
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		slog.Error("postgres", "err", "Unable to list authors", "details", err.Error())
		os.Exit(1)
	}
	return authors
}

func createOneAuthor(ctx context.Context, queries *db.Queries) db.Author {
	author, err := queries.CreateAuthor(ctx, db.CreateAuthorParams{
		Name: "Felipe Oliveira",
		Bio: pgtype.Text{
			String: "Felipe Oliveira é um desenvolvedor Golang talentoso com uma paixão por criar soluções eficientes e escaláveis",
			Valid:  true,
		},
	})

	if err != nil {
		slog.Error("postgres", "err", "Unable to create author", "details", err.Error())
	}

	return author
}

func main() {
	ctx := context.Background()
	conn := connectToBD(ctx)
	defer conn.Close()

	/*
		Para utilizar as queries geradas pelo sqlc, é necessário passar a conexão com o banco de dados.
		Depois disso, basta chamar as funções geradas pelo sqlc.
	*/
	queries := db.New(conn)

	authors := listAllAuthors(ctx, queries)
	slog.Info("postgres", "authors", authors)

	createdAuthor := createOneAuthor(ctx, queries)
	slog.Info("postgres", "author", fmt.Sprintf("%+v", createdAuthor))
}
