package transform

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
)

func ParseHTML(html string, data map[string]interface{}) string {
	doc := soup.HTMLParse(html)
	title := doc.FindAll("h1")
	for _, t := range title {
		name := strings.Replace(strings.Replace(t.Text(), " Ice Cream Recipe", "", -1), "Easy ", "", -1)
		data["name"] = name
	}
	data["description"] = doc.FindStrict("div", "class", "entry-content single-content").Find("p").FullText()
	// This may need to be cleaned a bit or just removed since it may not be accurate.
	data["serves"] = doc.FindStrict("span", "class", "mv-create-time-format mv-create-uppercase").Text()
	imageUrls := doc.FindAll("img")
	// Needs to be case insensitive.
	re := regexp.MustCompile(`jpeg|jpg|png`)
	for _, is := range imageUrls {
		imgSrc := is.Attrs()["src"]
		if re.MatchString(imgSrc) {
			// I need to study this copilot solution
			data["imageUrls"] = append(data["imageUrls"].([]string), imgSrc)
		}
	}
	ingredients := doc.Find("div", "class", "mv-create-ingredients").Find("ul").FindAll("li")
	for _, ing := range ingredients {
		ingText := strings.TrimSpace(ing.Text())
		data["ingredients"] = append(data["ingredients"].([]string), ingText)
	}
	instructions := doc.Find("div", "class", "mv-create-instructions").Find("ol").FindAll("li")
	for _, ins := range instructions {
		insText := strings.TrimSpace(ins.Text())
		data["instructions"] = append(data["instructions"].([]string), insText)
	}
	bytes, _ := json.Marshal(data)
	return string(bytes)
}
