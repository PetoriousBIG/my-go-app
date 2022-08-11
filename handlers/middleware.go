package handlers

import (
	"context"
	"net/http"

	"github.com/PetoriousBIG/my-go-app/data"
	"github.com/gorilla/mux"
)

func (c *countryData) GetMiddlewareValidateCountryFunc(dict *data.CountryDictionary) func(http.Handler) http.Handler {
	mw := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Add("Content-Type", "application/json")
			params := mux.Vars(r)
			countryCode := params["id"]
			header, ok := dict.Dict[countryCode]
			var ctx context.Context
			if ok {
				ctx = context.WithValue(r.Context(), "header", header)
				ctx = context.WithValue(ctx, "valid", true)

			} else {
				ctx = context.WithValue(r.Context(), "valid", false)
			}
			next.ServeHTTP(rw, r.WithContext(ctx))
		})
	}
	return mw
}
