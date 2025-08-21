package store

import "math/big"

type AccountNumeric struct {
	Amount *big.Int
	Exp    int32
}

type BaseAccount struct {
	ID           string
	MonthyIncome AccountNumeric
	Age          string
	Phone        string
	Email        string
	Category     string
	Balance      AccountNumeric
	CreatedAt    string
	UpdatedAt    string
	ClosedAt     string
}

type BusinessAccount struct {
	BaseAccount
	TradeName string
}

type PersonalAccount struct {
	BaseAccount
	FullName string
}
