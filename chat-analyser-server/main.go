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
	r.HandleFunc("/user/day/{author}", top.GetUserDayOfWeekHTTP).Methods("GET")
	r.HandleFunc("/user/hour/{author}", top.GetUserHourOfDayHTTP).Methods("GET")
	r.HandleFunc("/user/day", top.GetUserDayOfWeekHTTP).Methods("GET")
	r.HandleFunc("/user/hour", top.GetUserHourOfDayHTTP).Methods("GET")
	r.HandleFunc("/user/dist", top.GetUserDistChartHTTP).Methods("GET")
	r.HandleFunc("/date/range", top.GetDateRangeHTTP).Methods("GET")
	r.HandleFunc("/date/chat", top.GetChatLogHTTP).Methods("GET")
	r.HandleFunc("/date/dateChart", top.GetDateChartHTTP).Methods("GET")
	r.HandleFunc("/date/dateChart/keyword", top.GetDateChartKeywordHTTP).Methods("GET")
	http.ListenAndServe(":8000", r)

}
