package domain

type Finance struct {
	Base    string             `json:"base"`
	Rates   map[string]float64 `json:"rates"`
	Success bool               `json:"success"`
	Date    string             `json:"date"`
	Error   *FinanceError      `json:"error"`
}
type FinanceRequest struct {
	Base   string `json:"base"`
	ApiKey string `json:"apikey"`
}

type FinanceError struct {
	Code    int    `json:"code"`
	Message string `json:"type"`
}

type CurrencyCode struct {
	Name         string
	Alpha2Code   string
	CurrencyName string
	CurrencyCode string
}
