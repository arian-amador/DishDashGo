package main

import (
	"fmt"
	"log/slog"
	"os"

	"dish_dash_go/rss"
)

const (
	rssFeed = "https://icecreamfromscratch.com/feed/"
)

func main() {
	// NOTE: should we pass this logger around?
	handler := slog.NewTextHandler(os.Stderr, nil)
	logger := slog.New(handler)

	res, err := rss.Parse(rssFeed)
	if err != nil {
		logger.Error(err.Error())
	}

	// TODO: Parse rss feeds from res
	fmt.Println(res)
}
