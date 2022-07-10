package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/youcoulddoworse/vizrest/vizmse"
)

var client = vizmse.CreateClient()

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":3000", r)
	log.Println("Listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	s, err := vizmse.GetBase(client, "http://127.0.0.1:8580")
	if err != nil {
		log.Println(err)
		return
	}

	payload, err := json.Marshal(s)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)

}
