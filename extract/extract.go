package extract

import (
	"fmt"
	"io"
	"net/http"

	"github.com/asaskevich/govalidator"
)

func PageHTML(urlString string) (string, error) {
	// Validate URL
	u := govalidator.IsURL(urlString)
	if !u {
		errorMessage := fmt.Errorf("error - Invalid URL: %s", urlString)
		return "", errorMessage
	}
	// Retrieve HTML data
	resp, err := http.Get(string(urlString))
	fmt.Println(err)
	if err != nil {
		errorMessage := fmt.Errorf("error - %v", err)
		return "", errorMessage
	}
	// Validate HTML response
	if resp.StatusCode > 299 {
		errorMessage := fmt.Errorf("error - Non 200 status code %v: %s", resp.StatusCode, urlString)
		return "", errorMessage
	}
	// Extract HTML body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errorMessage := fmt.Errorf("error - %v", err)
		return "", errorMessage
	}
	return string(body), err
}
