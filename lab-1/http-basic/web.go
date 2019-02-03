package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func darMensaje(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Laboratory #1")
}

func handleDate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Laboratory at time: "+getDateString())
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func getDateString() string {
	now := time.Now()
	return now.Format(time.UnixDate)
}

func main() {
	http.HandleFunc("/", darMensaje)
	http.HandleFunc("/time", handleDate)
	http.HandleFunc("/index", index)
	err := http.ListenAndServe("localhost"+":"+"8080", nil)
	if err != nil {
		log.Fatal("error en el servidor : ", err)
		return
	}
}
