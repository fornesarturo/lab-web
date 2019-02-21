package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	cache "github.com/patrickmn/go-cache"
)

var (
	cach *cache.Cache
)

type People struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Gender    string   `json:"gender`
	Age       string   `json:"age"`
	EyeColor  string   `json:"eye_color"`
	HairColor string   `json:"hair_color"`
	Films     []string `json:"films"`
	Species   string   `json:"species"`
	URL       string   `json:url`
}

type PeopleErr struct {
	People []People
	Error  error
}

func cookieHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "Triki_Trakes", Value: "QueRicas", Expires: time.Now().Add(5 * time.Minute)})
	w.WriteHeader(200)
}

func fetchJSON(id string, c chan PeopleErr) {
	response, err := http.Get("http://ghibliapi.herokuapp.com/people/" + id)
	if err != nil {
		c <- PeopleErr{People: []People{}, Error: errors.New("Error in Ghibli request")}
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c <- PeopleErr{People: []People{}, Error: errors.New("Error reading body")}
		return
	}
	if id != "" {
		var a People
		json.Unmarshal(body, &a)
		c <- PeopleErr{People: []People{a}}
		return
	}
	var a []People
	json.Unmarshal(body, &a)
	c <- PeopleErr{People: a}
}

func readJSON(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path + r.URL.RawQuery
	var bytes []byte
	result, found := cach.Get(key)
	if found {
		log.Println("I have a byte!")
		bytes = result.([]byte)
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	} else {
		query := r.URL.Query()
		id := ""
		if query["id"] != nil {
			id = query["id"][0]
		}
		c := make(chan PeopleErr)
		go fetchJSON(id, c)
		peopleErr := <-c
		if peopleErr.Error != nil {
			log.Println("Error parsing JSON")
		}
		bytes, err := json.Marshal(peopleErr.People)
		if err != nil {
			log.Println("Error marshalling JSON")
			w.WriteHeader(500)
		}
		cach.Set(key, bytes, cache.NoExpiration)
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	}
}

func main() {
	cach = cache.New(5*time.Minute, 10*time.Minute)
	fmt.Println(time.Now())
	http.HandleFunc("/cookie", cookieHandler)
	http.HandleFunc("/readJSON", readJSON)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Something died here")
	}
}
