package handlers

// func Test_GetCountryDataWithValidCountry(t *testing.T) {
// 	l := log.New(os.Stdout, "UNIT TEST ", log.LstdFlags)
// 	c := NewCountryData(l)
// 	expectedAtAGlance := domain.AtAGlance{domain.CountryHeader{"TST", "Testland", "TS", 1, 0, 0}, data.Finance{}}

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
// 	expectedHeader := domain.ApiError{"country not found, {TST Testland TS 1 0 0}"}

// 	req := httptest.NewRequest("GET", "/", nil)
// 	ctx := req.Context()
// 	ctx = context.WithValue(ctx, "valid", false)
// 	ctx = context.WithValue(ctx, "header", arg)
// 	ctx = context.WithValue(ctx, "currency", nil)

// 	out := httptest.NewRecorder()
// 	c.GetCountryData(out, req.WithContext(ctx))

// 	actualStatus := out.Result().StatusCode
// 	var actualHeader data.ApiError
// 	util.FromJSON(&actualHeader, out.Body)

// 	assert.Equal(t, http.StatusNotFound, actualStatus)
// 	assert.Equal(t, expectedHeader, actualHeader)
// }
