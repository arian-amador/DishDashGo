package extract

import (
	"io"
	"log"
	"net/http"
)

func PageHTML(url string) string {
	resp, err := http.Get(string(url))
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}
