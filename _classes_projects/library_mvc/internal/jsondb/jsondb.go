package jsondb

type Table int

const (
	USERS Table = iota
	BOOKS
	LOANS
)

type JsonDB struct{}

func (j *JsonDB) get(table Table) {

}
