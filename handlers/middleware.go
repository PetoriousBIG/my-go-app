package handlers

import (
	"context"
	"net/http"

	"github.com/PetoriousBIG/my-go-app/data"
	"github.com/gorilla/mux"
)

func (c *countryData) GetMiddlewareValidateCountryFunc(headers map[string]data.CountryHeader, currencies map[string]data.CurrencyCode) func(http.Handler) http.Handler {
	mw := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			var currency data.CurrencyCode
			rw.Header().Add("Content-Type", "application/json")
			params := mux.Vars(r)
			countryCode := params["id"]
			header, ok := headers[countryCode]

			var ctx context.Context
			ctx = context.WithValue(r.Context(), "header", header)

			if ok {
				currency, ok = currencies[header.Alpha2Code]
				ctx = context.WithValue(ctx, "currency", currency)
			}

			if ok {
				ctx = context.WithValue(ctx, "valid", true)
			} else {
				ctx = context.WithValue(r.Context(), "valid", false)
			}

			next.ServeHTTP(rw, r.WithContext(ctx))
		})
	}
}
