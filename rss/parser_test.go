package rss

import "testing"

func TestPageNotFoundError(t *testing.T) {
	want := "error: page not found"
	_, got := Parse("https://icecreamfromscratch.com/feed-errorrrrr/")

	if got == nil {
		t.Errorf("Did not get error on invalid url")
	} else if got.Error() != want {
		t.Errorf("Did not get the expected error message %q, but %q", want, got.Error())
	}
}
