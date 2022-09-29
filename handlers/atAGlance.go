package handlers

// type Getter func([]interface{}) interface{}
// type countryData struct {
// 	l *log.Logger
// 	g Getter
// }

// func NewCountryData(l *log.Logger, g Getter) *countryData {
// 	return &countryData{l, g}
// }

// func (c *countryData) GetCountryData(rw http.ResponseWriter, r *http.Request) {
// 	valid := r.Context().Value("valid").(bool)
// 	header := r.Context().Value("header").(data.CountryHeader)

// 	rw.Header().Add("Content-Type", "application/json")

// 	var response any

// 	if valid {
// 		c.l.Println("[DEBUG] getting country data", header)
// 		rw.WriteHeader(http.StatusOK)
// 		currency := r.Context().Value("currency").(data.CurrencyCode)
// 		response = c.g([]interface{header currency})
// 	} else {
// 		c.l.Println("[DEBUG] country not found", header)
// 		rw.WriteHeader(http.StatusNotFound)
// 		response = data.ApiError{fmt.Sprintf("country not found, %v", header)}
// 	}

// 	util.ToJSON(response, rw)
// }

// func apiMashup(ch data.CountryHeader, cc data.CurrencyCode) *data.AtAGlance {

// 	base := cc.CurrencyCode
// 	finance := clients.NewFinacne(base)
// 	exchangeRates := finance.GET()

// 	output := data.AtAGlance{ch, exchangeRates}

// 	return &output
// }
