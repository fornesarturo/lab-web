package main

import (
	"io"
	"log"
	"net/url"
	"os"

	"github.com/knakk/rdf"
	"github.com/piprate/json-gold/ld"
)

func main() {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.ProcessingMode = ld.JsonLd_1_1
	options.Format = "application/n-quads"
	options.Algorithm = "URDNA2015"

	myURL := "https://kgsearch.googleapis.com/v1/entities:search"
	query := url.Values{}
	query.Add("query", "taylor swift")
	query.Add("key", "AIzaXXXXXXX")
	query.Add("limit", "50")
	query.Add("indent", "True")

	finalURL := myURL + "?" + query.Encode()

	normalized, err := proc.Normalize(finalURL, options)

	// fmt.Println(normalized)

	if err != nil {
		log.Println(normalized, err)
	}

	f, err := os.Open("triplets.ttl")
	if err != nil {
		log.Println(err)
	}
	dec := rdf.NewTripleDecoder(f, rdf.Turtle)
	for triple, err := dec.Decode(); err != io.EOF; triple, err = dec.Decode() {
		if triple.Pred.String() == "http://schema.org/articleBody" {
			log.Println(triple.Obj)
		}
	}
}
