package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// handling request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[Hello World]")
		d, err := io.ReadAll(r.Body)
		//log.Printf("Data %s\n", d)
		if err != nil {
			http.Error(w, "Ooops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello %s\n", d)
	})
	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[Goodbye World]")
	})
	http.ListenAndServe(":9090", nil)
}
