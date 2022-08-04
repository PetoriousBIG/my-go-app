package handlers

import (
	"context"
	"encoding/json"
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
			if !ok {
				c.l.Println("[ERROR] country not found")

				rw.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(rw).Encode(map[string]interface{}{
					"message":  "resource not found",
					"endpoint": countryCode,
				})
				return
			}

			ctx := context.WithValue(r.Context(), "header", header)
			next.ServeHTTP(rw, r.WithContext(ctx))

		})
	}
	return mw
}
