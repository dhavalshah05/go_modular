package models

type AddTransactionRequest struct {
	Credit   int    `json:"credit"`
	Debit    int    `json:"debit"`
	Category string `json:"category"`
	Summary  string `json:"summary"`
}

type Transaction struct {
	Id       int    `json:"id"`
	Credit   int    `json:"credit"`
	Debit    int    `json:"debit"`
	Category string `json:"category"`
	Summary  string `json:"summary"`
}
