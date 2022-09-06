package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type finance struct {
	BaseCurrency string
}

func NewFinacne(bc string) *finance {
	return &finance{bc}
}

func (f *finance) GET() map[string]interface{} {
	client := http.Client{}
	u := "https://api.exchangerate.host/latest"
	request, err := http.NewRequest("GET", u, nil)
	values := request.URL.Query()
	values.Add("base", f.BaseCurrency)
	request.URL.RawQuery = values.Encode()
	request.Header.Add("Accept", "application/json")

	fmt.Printf("%v", request.URL)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result
}
