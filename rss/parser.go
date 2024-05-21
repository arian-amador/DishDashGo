package rss

import (
	"fmt"
	"io"
	"net/http"
)

// TODO:: a feed or parser should be a custom type with an interface and methods to read/parse/whatever

// TODO: We should def not return a string here and something else
func Parse(rssFeed string) (string, error) {
	resp, err := http.Get(rssFeed)
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
