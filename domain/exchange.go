package domain

type Exchange struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
	Date  string             `json:"date"`
}
type ExchangeResponse struct {
	Base    string             `json:"base"`
	Rates   map[string]float64 `json:"rates"`
	Success bool               `json:"success"`
	Date    string             `json:"date"`
	Error   ExchangeError      `json:"error"`
}
type ExchangeRequest struct {
	Base   string `json:"base"`
	ApiKey string `json:"apikey"`
}
type ExchangeError struct {
	Code    int    `json:"code"`
	Message string `json:"type"`
}
