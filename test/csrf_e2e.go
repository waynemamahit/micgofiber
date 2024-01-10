package test

import (
	"micgofiber/lib"
	"net/http"
	"strings"
)

type CsrfE2E struct {
	app       *lib.AppConfig
	token     string
	setCookie string
}

func NewCsrfE2E(app *lib.AppConfig) *CsrfE2E {
	return &CsrfE2E{
		app:       app,
		token:     "",
		setCookie: "",
	}
}

func (e2e *CsrfE2E) Request(url string, method string, body any) (*http.Response, error) {
	resp, err := lib.RequestTest(e2e.app, e2e.app.ApiV1+url, method, body, [][2]string{
		{e2e.app.CsrfHeader, e2e.token},
		{"Cookie", e2e.setCookie},
	})

	if method == "GET" {
		var setCookie []string
		for _, cookie := range resp.Cookies() {
			if cookie.Name == e2e.app.CsrfCookie {
				e2e.token = cookie.Value
			}
			setCookie = append(setCookie, cookie.Name+"="+cookie.Value)
		}
		e2e.setCookie = strings.Join(setCookie, "; ")
	}

	return resp, err
}

func (e2e *CsrfE2E) FormRequest(url string, method string, fieldname string, filename string, dto any) (*http.Response, error) {
	return lib.FormDataRequestTest(e2e.app, e2e.app.ApiV1+url, method, fieldname, filename, dto, [][2]string{
		{e2e.app.CsrfHeader, e2e.token},
		{"Cookie", e2e.setCookie},
	})
}
