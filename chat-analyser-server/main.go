package main

import (
	"net/http"
	"search-analysis/res/top"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/top", top.GetTotalTop100KeywordsHTTP).Methods("GET")

	http.ListenAndServe(":8000", r)

}
