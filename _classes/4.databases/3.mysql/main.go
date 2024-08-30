package main

/*
	Para se conectar a um banco de dados MySQL, é necessário instalar o pacote go-sql-driver/mysql (mais usado).
	Ele implementa a interface da standard library database/sql.

	(Sem segredo, igual os outros drivers que implementam a interface database/sql)
*/

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	// Semelhante aos outros pacotes de drivers
	_ "github.com/go-sql-driver/mysql"
)

func connectToBD() *sql.DB {
	// Semelhante aos outros pacotes de drivers
	db, err := sql.Open("mysql", "user:password@/main")
	if err != nil {
		panic(err)
	}
	// Configurações de conexão
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	slog.Info("mysql", "connectToBD", "Successfully connected to the database")
	return db
}

func createTables(db *sql.DB) {
	query := "create table if not exists users (id bigint auto_increment primary key, name varchar(255), age int);"
	if _, err := db.Exec(query); err != nil {
		slog.Error("mysql", "createTables", err.Error())
	} else {
		slog.Info("mysql", "createTables", "Table users created successfully")
	}
}

func insertUser(db *sql.DB, name string, age int) {
	query := "insert into users (name, age) values (?, ?);"
	if _, err := db.Exec(query, name, age); err != nil {
		slog.Error("mysql", "insertUser", err.Error())
	} else {
		slog.Info("mysql", "insertUser", fmt.Sprintf("User %s with age %d inserted successfully", name, age))
	}
}

type user struct {
	id   int
	name string
	age  int
}

func selectAllUsers(db *sql.DB) []user {
	res := []user{}

	query := "select * from users;"
	/*
		O método Query retorna um objeto do tipo Rows, que é um conjunto de linhas retornadas por uma consulta.
	*/
	rows, err := db.Query(query)
	if err != nil {
		slog.Error("mysql", "selectAllUsers", err.Error())
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
		var u user
		if err := rows.Scan(&u.id, &u.name, &u.age); err != nil {
			slog.Error("mysql", "selectAllUsers", err.Error())
		}
		res = append(res, u)
	}

	return res
}

func main() {
	db := connectToBD()

	createTables(db)
	insertUser(db, "John", 30)

	users := selectAllUsers(db)
	slog.Info("mysql", "users", fmt.Sprintf("%+v", users))
}
