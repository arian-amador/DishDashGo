package main

import (
	"dish_dash_go/extract"
	"dish_dash_go/transform"
	"fmt"
	"sync"
)

/*
 I need to figure out how to deal with and recover from panics. For now I
 need to write a helper funtion that sanitizes the input to avoid panics
 where possible.
*/

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
		"ttps://icecreamfromscratch.com/cornbread-ice-creams/",
		"https://icecreamfromscratch.com/cornbread-ice-creams/",
		"https://icecreamfromscratch.com/cornbread-ice-cream/",
		"https://icecreamfromscratch.com/lime-ice-cream/",
		"https://icecreamfromscratch.com/honey-lavender-ice-cream/",
		"https://icecreamfromscratch.com/blueberry-cheesecake-ice-cream/",
	}

	waitgroup.Add(len(urls))
	for _, url := range urls {
		go func(u string) {
			defer waitgroup.Done()
			html, err := extract.PageHTML(u)
			if err != nil {
				fmt.Println(err)
				return
			}
			// jsonString := transform.ParseHTML(html, data)
			transform.ParseHTML(html, data)
		}(url)
	}
	waitgroup.Wait()
}
