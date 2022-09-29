package handlers

import (
	"testing"

	"github.com/PetoriousBIG/my-go-app/domain"
	"github.com/PetoriousBIG/my-go-app/providers"
)

var (
	financeProviderFunc func(r domain.FinanceRequest) *domain.Finance
)

type financeProviderMock struct{}

//We are mocking the provider method "GetFinance"
func (fp *financeProviderMock) GetFinance(r domain.FinanceRequest) *domain.Finance {
	return financeProviderFunc(r)
}

func TestGetAtAGlanceWithValidCountry(t *testing.T) {
	financeProviderFunc = func(r domain.FinanceRequest) *domain.Finance {
		return &domain.Finance{}
	}
	providers.FinanceProvider = &financeProviderMock{} // mock applied

}

// func Test_GetCountryDataWithValidCountry(t *testing.T) {
// 	l := log.New(os.Stdout, "UNIT TEST ", log.LstdFlags)
// 	c := NewCountryData(l)
// 	expectedAtAGlance := domain.AtAGlance{domain.CountryHeader{"TST", "Testland", "TS", 1, 0, 0}, domain.Finance{}}

// 	req := httptest.NewRequest("GET", "/", nil)
// 	ctx := req.Context()
// 	ctx = context.WithValue(ctx, "valid", true)
// 	ctx = context.WithValue(ctx, "header", expectedAtAGlance.CountryHeader)
// 	ctx = context.WithValue(ctx, "currency", domain.CurrencyCode{"Testland", "TS", "Testland Dollars", "TSD"})

// 	out := httptest.NewRecorder()
// 	c.GetCountryData(out, req.WithContext(ctx))

// 	actualStatus := out.Result().StatusCode
// 	var actualAtAGlance domain.AtAGlance
// 	util.FromJSON(&actualAtAGlance, out.Body)

// 	assert.Equal(t, http.StatusOK, actualStatus)
// 	assert.Equal(t, expectedAtAGlance, actualAtAGlance)

// }

// func Test_GetCountryDataWithInvalidCountry(t *testing.T) {
// 	l := log.New(os.Stdout, "UNIT TEST ", log.LstdFlags)
// 	c := NewCountryData(l)
// 	arg := domain.CountryHeader{"TST", "Testland", "TS", 1, 0, 0}
// 	expectedHeader := domain.AtAGlanceError{}

// 	req := httptest.NewRequest("GET", "/", nil)
// 	ctx := req.Context()
// 	ctx = context.WithValue(ctx, "valid", false)
// 	ctx = context.WithValue(ctx, "header", arg)
// 	ctx = context.WithValue(ctx, "currency", nil)

// 	out := httptest.NewRecorder()
// 	c.GetCountryData(out, req.WithContext(ctx))

// 	actualStatus := out.Result().StatusCode
// 	var actualHeader domain.AtAGlanceError
// 	util.FromJSON(&actualHeader, out.Body)

// 	assert.Equal(t, http.StatusNotFound, actualStatus)
// 	assert.Equal(t, expectedHeader, actualHeader)
// }
