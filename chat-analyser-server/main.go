package main

import (
	"chat-analyser/res/top"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/top", top.GetTotalTop100KeywordsHTTP).Methods("GET")
	r.HandleFunc("/user", top.GetTotalUsersHTTP).Methods("GET")
	r.HandleFunc("/user/keywords/{author}", top.GetUserKeywordsHTTP).Methods("GET")
	http.ListenAndServe(":8000", r)

}
