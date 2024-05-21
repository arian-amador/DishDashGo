package rss

import (
	"fmt"
	"io"
	"net/http"
)

type parser struct {
	url string
}

// New method that returns a parser struct
func NewParser(url string) parser {
	return(parser{url: url})
} 

// Method to read and return an rss feed
func (p *parser) Read() (string, error) {
	resp, err := http.Get(p.url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: page not found")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
