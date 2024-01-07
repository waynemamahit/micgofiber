package lib

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func RequestTest(app *AppConfig, prefix string, method string, body any, headers [][2]string) (*http.Response, error) {
	// Create a new http request with the route from the test case
	data, errData := json.Marshal(body)
	if errData != nil {
		return nil, errData
	}
	req := httptest.NewRequest(method, app.ApiV1+prefix, bytes.NewReader(data))
	for _, header := range headers {
		if len(header) == 2 {
			req.Header.Set(header[0], header[1])
		}
	}

	// Perform the request plain with the app,
	// the second argument is a request latency
	// (set to -1 for no latency)
	resp, err := app.App.Test(req, -1)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
