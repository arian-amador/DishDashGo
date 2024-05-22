package main

import (
	"dish_dash_go/extract"
	"dish_dash_go/transform"
	"sync"
)

func main() {
	data := map[string]interface{}{
		"name":         "",
		"imageUrls":    []string{},
		"description":  "",
		"serves":       0,
		"ingredients":  []string{},
		"instructions": []string{},
	}

	var waitgroup sync.WaitGroup
	// RSS will feed this. May connect Amandor's code later.
	urls := []string{
		"https://icecreamfromscratch.com/cornbread-ice-cream/",
		"https://icecreamfromscratch.com/lime-ice-cream/",
		"https://icecreamfromscratch.com/honey-lavender-ice-cream/",
		"https://icecreamfromscratch.com/blueberry-cheesecake-ice-cream/",
	}

	waitgroup.Add(len(urls))
	for _, url := range urls {
		go func(u string) {
			defer waitgroup.Done()
			html := extract.PageHTML(u)
			transform.ParseHTML(html, data)
		}(url)
	}
	waitgroup.Wait()
}
