package main

import (
	"log"
	"net/url"

	"github.com/piprate/json-gold/ld"
)

func main() {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	myURL := "https://kgsearch.googleapis.com/v1/entities:search"
	query := url.Values{}
	query.Add("query", "taylor swift")
	query.Add("key", "AIzaXXXXXXX")
	query.Add("limit", "50")
	query.Add("indent", "True")

	finalURL := myURL + "?" + query.Encode()

	log.Println(finalURL)

	expanded, err := proc.Expand(finalURL, options)

	if err != nil {
		log.Println(expanded, err)
	}

	listElements := expanded[0].(map[string]interface{})["http://schema.org/itemListElement"].([]interface{})
	oneResult := listElements[len(listElements)-1].(map[string]interface{})["http://schema.org/result"].([]interface{})[0].(map[string]interface{})
	oneResultDescription := oneResult["http://schema.org/description"].([]interface{})[0].(map[string]interface{})["@value"].(string)
	log.Println("Worst possible description: ", oneResultDescription)
}
