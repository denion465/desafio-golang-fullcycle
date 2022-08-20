package models

type Transaction struct {
	ID           string
	From_account string
	To_account   string
	Amount       float64
}
