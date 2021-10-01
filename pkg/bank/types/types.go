package types

// тип денег
type Money int64

// тип валюта
type Currency string

//подтипы валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type PAN string

//структура карты
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
	MinBalance Money
}

type Payment struct {
	ID     int
	Amount Money
}
