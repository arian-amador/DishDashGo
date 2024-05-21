package rss

import (
	"testing"
)

func TestNewParser(t *testing.T) {
	url := "http://url"
	parser := NewParser(url)

	if parser.url != url {
		t.Errorf("Error creating a new parser with url %q", url)
	}
}

func TestPageNotFoundError(t *testing.T) {
	p := NewParser("https://icecreamfromscratch.com/feed-errorrrrr/")
	want := "error: page not found"
	_, got := p.Read()

	if got == nil {
		t.Errorf("Did not get error on invalid url")
	} else if got.Error() != want {
		t.Errorf("Did not get the expected error message %q, but %q", want, got.Error())
	}
}
