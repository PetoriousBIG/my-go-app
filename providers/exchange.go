package providers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/PetoriousBIG/my-go-app/clients"
	"github.com/PetoriousBIG/my-go-app/domain"
)

const (
	financeURL = "https://api.apilayer.com/fixer/latest"
)

type exchangeProvider struct{}

type iExchangeProvider interface {
	GetExchange(request domain.ExchangeRequest) (*domain.Exchange, *domain.ExchangeError)
}

var (
	ExchangeProvider iExchangeProvider = &exchangeProvider{}
)

func (p *exchangeProvider) GetExchange(request domain.ExchangeRequest) (*domain.Exchange, *domain.ExchangeError) {
	r, err := http.NewRequest("GET", financeURL, nil)
	if err != nil {
		return nil, &domain.ExchangeError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	vals := r.URL.Query()
	vals.Add("base", request.Base)
	vals.Add("apikey", request.ApiKey)
	r.URL.RawQuery = vals.Encode()
	resp, err := clients.ClientStruct.Get(r)
	if err != nil {
		return nil, &domain.ExchangeError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, &domain.ExchangeError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	defer resp.Body.Close()

	var exchangeResponse domain.ExchangeResponse
	if err := json.Unmarshal(bytes, &exchangeResponse); err != nil {
		return nil, &domain.ExchangeError{
			Code:    exchangeResponse.Error.Code,
			Message: "invalid json response body",
		}
	}

	if resp.StatusCode > 299 || !exchangeResponse.Success {
		return nil, &domain.ExchangeError{
			Code:    exchangeResponse.Error.Code,
			Message: exchangeResponse.Error.Message,
		}
	}

	result := domain.Exchange{
		Base:  exchangeResponse.Base,
		Rates: exchangeResponse.Rates,
		Date:  exchangeResponse.Date,
	}
	return &result, nil
}
