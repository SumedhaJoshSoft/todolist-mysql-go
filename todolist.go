package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	status := 200
	log.Println("API Health is OK")
	respBytes, err := json.Marshal(`{"alive":true}`)
	if err != nil {
		log.Println(err)
		status = http.StatusInternalServerError
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(respBytes)
}

func main() {
	log.Println("Starting Todolist API server")
	router := mux.NewRouter()
	router.HandleFunc("/healthz", Healthz).Methods("GET")
	http.ListenAndServe(":8000", router)
}
