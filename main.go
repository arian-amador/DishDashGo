package main

import (
	"fmt"
	"log/slog"
	"os"

	"dish_dash_go/rss"
)

const (
	rssFeedUrl = "https://icecreamfromscratch.com/feed/"
)

func main() {
	// NOTE: should we pass this logger around?
	handler := slog.NewTextHandler(os.Stderr, nil)
	logger := slog.New(handler)

	p := rss.NewParser(rssFeedUrl)
	res, err := p.Read()
	if err != nil {
		logger.Error(err.Error())
	}

	// TODO: Parse rss feeds from res
	fmt.Println(res)
}
