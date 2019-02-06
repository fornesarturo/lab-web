package main

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	sites []string = []string{
		"http://www.google.com",
		"http://www.reddit.com",
		"http://www.facebook.com",
		"http://golang.org",
		"https://doc.rust-lang.org/",
		"https://isitchristmas.com/",
		"https://stackoverflow.com",
		"https://medium.com",
		"https://hackertyper.net/",
		"https://www.9gag.com/",
		"https://github.com",
		"https://vuejs.org/",
		"https://netflix.com/",
		"https://youtube.com",
		"https://dev.to/",
		"https://www.invisionapp.com/",
		"https://www.linkedin.com/",
		"https://theuselessweb.com/",
		"https://cant-not-tweet-this.com/",
		"https://twitter.com/",
		"https://weirdorconfusing.com/",
		"https://amazon.com",
		"https://heeeeeeeey.com/",
		"https://hooooooooo.com/",
		"http://tinytuba.com/",
	}
)

func getHeader(site string, ch chan map[string][]string, w *sync.WaitGroup) {
	res, err := http.Get(site)
	if err == nil {
		fmt.Println("site:", site)
		w.Done()
		ch <- res.Header
	} else {
		fmt.Println("Error:", err)
		w.Done()
	}
}

func main() {
	var w sync.WaitGroup
	w.Add(len(sites))
	ch := make(chan map[string][]string)
	for _, site := range sites {
		go getHeader(site, ch, &w)
	}
	w.Wait()
	for i := 0; i < len(sites); i++ {
		head := <-ch
		fmt.Println(i, "\n", head)
	}
}
