package lib

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
)

func RequestTest(
	app *AppConfig,
	url string,
	method string,
	body any,
	headers [][2]string,
) (*http.Response, error) {
	// Create a new http request with the route from the test case
	data, errData := json.Marshal(body)
	if errData != nil {
		return nil, errData
	}
	req := httptest.NewRequest(method, url, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
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

func FormDataRequestTest(
	app *AppConfig,
	url string,
	method string,
	fieldname string,
	filename string,
	dto any,
	headers [][2]string,
) (*http.Response, error) {
	// Open the file for reading
	file, err := os.Open("../test/mock/file/" + filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new temporary file
	tempFile, err := os.CreateTemp("", strings.Join(strings.Split(filename, "."), "-*."))
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()

	// Copy the content of the original file to the temporary file
	_, err = io.Copy(tempFile, file)
	if err != nil {
		return nil, err
	}

	// Create a new file to upload
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// Create a new form file field
	part, err := writer.CreateFormFile(fieldname, filepath.Base(tempFile.Name()))
	if err != nil {
		return nil, err
	}

	// Open the temporary file and copy its content to the form file field
	tempFile.Seek(0, 0)
	_, err = io.Copy(part, tempFile)
	if err != nil {
		return nil, err
	}

	// Set additional fields
	jsonStr, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}
	var mapJson map[string]string
	err = json.Unmarshal(jsonStr, &mapJson)
	if err != nil {
		return nil, err
	}
	for key, v := range mapJson {
		_ = writer.WriteField(key, v)
	}

	// Close the multipart writer
	writer.Close()

	// Create a request to the specified URL
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	for _, header := range headers {
		if len(header) == 2 {
			req.Header.Set(header[0], header[1])
		}
	}

	// Perform the request
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
