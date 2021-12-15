package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/api/", Num)
	log.Println("listening on port :9999")
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatalln(err)
	}
}

func Num(w http.ResponseWriter, r *http.Request) {
	n := r.URL.Query().Get("n")
	fmt.Fprintf(w, "request num: %v", n)
}
