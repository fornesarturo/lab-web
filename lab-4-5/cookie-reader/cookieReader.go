package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	r, err := http.Get("http://localhost:8080/cookie")
	if err != nil {
		log.Fatal("I'm the one that's dying")
	}
	cookies := r.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "Triki_Trakes" {
			fmt.Println("Cookie!: ", cookie)
		}
	}
	fmt.Println(r)
}
