package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/PetoriousBIG/my-go-app/data"
	"github.com/gorilla/mux"
)

// MiddlewareValidateCountry validates the product in the request and call next if of
func (c *countryData) MiddlewareValidateCountry(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		params := mux.Vars(r)
		countryCode := params["id"]

		_, ok := data.CountryDictionary[countryCode]
		if !ok {
			c.l.Println("[ERROR] country not found")

			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(map[string]interface{}{
				"message":  "resource not found",
				"endpoint": countryCode,
			})
			return
		}

		next.ServeHTTP(rw, r)
	})
}
