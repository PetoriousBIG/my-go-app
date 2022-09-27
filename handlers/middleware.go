package handlers

import (
	"context"
	"net/http"

	"github.com/PetoriousBIG/my-go-app/domain"
	"github.com/gorilla/mux"
)

func (c *countryData) GetMiddlewareValidateCountryFunc(headers map[string]domain.CountryHeader, currencies map[string]domain.CurrencyCode) func(http.Handler) http.Handler {
	mw := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			var currency domain.CurrencyCode
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
	return mw
}
