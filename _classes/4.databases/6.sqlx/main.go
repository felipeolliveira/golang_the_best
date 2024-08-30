package main

import (
	"database/sql"
	"log/slog"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

/*
	O SQLx é essencialmente a mesma coisa que o SQLc, porém é feita em tempo de execução, já o SQLc é feito em tempo de compilação, pois compila o código antes de executar

	O SQLx faz isso atravez de reflection, o que pode ser mais lento que o SQLc, porém é mais flexível, pois usa esse recurso para mapear os dados do banco de dados para a struct
*/

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country string
	// O sql.NullString é um tipo que pode ser nulo, então se o campo city for nulo, o valor será nil
	// este tipo está disponível no pacote database/sql
	City    sql.NullString
	TelCode int
}

/*
Para conectar ao db, segue semelhante a outros drivers, porém a conexão é feita internamente com o sqlx.Connect que nesse caso usa o pq
*/
func connectToDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=user dbname=sqlx password=password port=5450 sslmode=disable")

	if err != nil {
		slog.Error("sqlx", "error on connect", err)
		panic(err)
	}

	slog.Info("sqlx", "msg", "connected to database")
	return db
}

func createSchemas(db *sqlx.DB) {
	schema := `
		DROP TABLE IF EXISTS person;
		DROP TABLE IF EXISTS place;
	
		CREATE TABLE person (
			first_name text,
			last_name text,
			email text
		);

		CREATE TABLE place (
			country text,
			city text NULL,
			telcode integer
		)`

	db.MustExec(schema)
}

func createInitialPersons(db *sqlx.DB) {
	// O TX é uma transaction, que é uma forma de agrupar várias queries em uma única transação, assim se uma falhar, todas falham.
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
	tx.MustExec("INSERT INTO place (country, city, telcode) VALUES ($1, $2, $3)", "United States", "New York", "1")
	tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Hong Kong", "852")
	tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Singapore", "65")

	// Named queries são queries que usam o nome dos campos da struct para mapear os valores, ao invés de usar o $1, $2, $3, etc
	tx.NamedExec(
		"INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)",
		&Person{"Jane", "Citizen", "jane.citzen@example.com"},
	)

	// Commit é usado para finalizar a transaction
	tx.Commit()
	slog.Info("sqlx", "msg", "created initial persons")
}

func selectPeople(db *sqlx.DB) []Person {
	people := []Person{}

	/*
		Passando a referência de um slice de Person, o sqlx irá mapear os valores do banco de dados para a struct Person
	*/
	if err := db.Select(&people, "select * from person order by first_name asc"); err != nil {
		slog.Error("sqlx", "error on get people", err)
		panic(err)
	}

	return people
}

func getAPersonByName(db *sqlx.DB, name string) Person {
	jason := Person{}
	err := db.Get(&jason, "SELECT * FROM person WHERE first_name=$1", name)
	if err != nil {
		slog.Error("sqlx", "error on get a person", err)
		panic(err)
	}
	return jason
}

func selectPlaces(db *sqlx.DB) []Place {
	places := []Place{}

	/*
		Como o Place tem um campo que pode ser nulo, é necessário usar o sql.NullString, que é um tipo que pode ser nulo, para mapear o campo city
	*/
	err := db.Select(&places, "select * from place order by telcode asc")
	if err != nil {
		slog.Error("sqlx", "error on get places", err)
		panic(err)
	}
	return places
}

func usingNamedQuery(db *sqlx.DB) {
	// É possível usar named queries para fazer selects com mapas genericos
	people := []Person{}
	rowsFromMapSelect, err := db.NamedQuery(`SELECT * FROM person WHERE first_name=:fn`, map[string]interface{}{"fn": "Bin"})
	if err != nil {
		slog.Error("sqlx", "error named queries", err)
		panic(err)
	}
	defer rowsFromMapSelect.Close()
	for rowsFromMapSelect.Next() {
		person := Person{}
		// Usando StructScan para mapear os valores do banco de dados para a struct Person
		err := rowsFromMapSelect.StructScan(&person)
		if err != nil {
			slog.Error("sqlx", "error named queries", err)
			panic(err)
		}
		people = append(people, person)
	}
	slog.Info("sqlx", "usingNamedQuery map", people)

	// Named queries também podem ser usadas com structs
	people = []Person{}
	jason := Person{FirstName: "Jason"}
	rowsFromStructSelect, err := db.NamedQuery(`SELECT * FROM person WHERE first_name=:first_name`, jason)
	if err != nil {
		slog.Error("sqlx", "error named queries", err)
		panic(err)
	}
	defer rowsFromStructSelect.Close()
	for rowsFromStructSelect.Next() {
		person := Person{}
		err := rowsFromStructSelect.StructScan(&person)
		if err != nil {
			slog.Error("sqlx", "error named queries", err)
			panic(err)
		}
		people = append(people, person)
	}

	slog.Info("sqlx", "usingNamedQuery struct", people)
}

func usingNamedExecs(db *sqlx.DB) {
	// É possível usar mapas genericos para fazer um insert com named queries
	_, err := db.NamedExec(`INSERT INTO person (first_name,last_name,email) VALUES (:first,:last,:email)`,
		map[string]interface{}{
			"first": "Bin",
			"last":  "Smuth",
			"email": "bensmith@allblacks.nz",
		})
	if err != nil {
		slog.Error("sqlx", "error named queries", err)
		panic(err)
	}

	/*
		Batch Inserts
		Com o sqlx é possível fazer batch inserts, que é inserir vários registros de uma vez usando NamedExec()
		Para isso é necessário passar um slice de structs ou um slice de mapas genericos que representam os registros a serem inseridos
	*/

	// batch insert com structs
	personStructs := []Person{
		{FirstName: "Ardie", LastName: "Savea", Email: "asavea@ab.co.nz"},
		{FirstName: "Sonny Bill", LastName: "Williams", Email: "sbw@ab.co.nz"},
		{FirstName: "Ngani", LastName: "Laumape", Email: "nlaumape@ab.co.nz"},
	}
	_, err = db.NamedExec(`INSERT INTO person (first_name, last_name, email)
	VALUES (:first_name, :last_name, :email)`, personStructs)
	if err != nil {
		slog.Error("sqlx", "error named queries", err)
		panic(err)
	}

	// batch insert com mapas genericos
	personMaps := []map[string]interface{}{
		{"first_name": "Ardie", "last_name": "Savea", "email": "asavea@ab.co.nz"},
		{"first_name": "Sonny Bill", "last_name": "Williams", "email": "sbw@ab.co.nz"},
		{"first_name": "Ngani", "last_name": "Laumape", "email": "nlaumape@ab.co.nz"},
	}
	_, err = db.NamedExec(`INSERT INTO person (first_name, last_name, email)
	VALUES (:first_name, :last_name, :email)`, personMaps)

	if err != nil {
		slog.Error("sqlx", "error named queries", err)
		panic(err)
	}

	slog.Info("sqlx", "usingNamedExec", selectPeople(db))
}

func main() {
	db := connectToDB()
	defer db.Close()

	createSchemas(db)
	createInitialPersons(db)

	people := selectPeople(db)
	slog.Info("sqlx", "people", people)

	jason := getAPersonByName(db, "Jason")
	slog.Info("sqlx", "jason", jason)

	places := selectPlaces(db)
	slog.Info("sqlx", "places", places)

	usingNamedQuery(db)
	usingNamedExecs(db)
}
