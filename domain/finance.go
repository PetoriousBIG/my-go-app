package domain

type Finance struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

type FinanceRequest struct {
	Base string `json:"base"`
}

type CurrencyCode struct {
	Name         string
	Alpha2Code   string
	CurrencyName string
	CurrencyCode string
}
