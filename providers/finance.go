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

type financeProvider struct{}

type financeService interface {
	GetFinance(request domain.FinanceRequest) *domain.Finance
}

var (
	FinanceProvider financeService = &financeProvider{}
)

func (p *financeProvider) GetFinance(request domain.FinanceRequest) *domain.Finance {
	r, err := http.NewRequest("GET", financeURL, nil)
	if err != nil {
		return &domain.Finance{
			Date:    "",
			Base:    "",
			Rates:   nil,
			Success: false,
			Error: &domain.FinanceError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}
	}

	vals := r.URL.Query()
	vals.Add("base", request.Base)
	vals.Add("apikey", request.ApiKey)
	r.URL.RawQuery = vals.Encode()
	resp, err := clients.ClientStruct.Get(r)
	if err != nil {
		return &domain.Finance{
			Date:    "",
			Base:    "",
			Rates:   nil,
			Success: false,
			Error: &domain.FinanceError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		}
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &domain.Finance{
			Date:    "",
			Base:    "",
			Rates:   nil,
			Success: false,
			Error: &domain.FinanceError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return &domain.Finance{
			Date:    "",
			Base:    "",
			Rates:   nil,
			Success: false,
			Error: &domain.FinanceError{
				Code:    resp.StatusCode,
				Message: err.Error(),
			},
		}
	}

	var result domain.Finance
	if err := json.Unmarshal(bytes, &result); err != nil {
		return &domain.Finance{
			Date:    "",
			Base:    "",
			Rates:   nil,
			Success: false,
			Error: &domain.FinanceError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}
	}
	return &result
}
