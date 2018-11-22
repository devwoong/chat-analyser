package top

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetTotalUsersHTTP is
func GetTotalUsersHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	query := `
	
	{
		"size": 0, 
		"query": {
			"match_all": {}
		},
		"aggs": {
			"author": {
				"terms": {
					"field": "author",
					"size": 100
				}
			}
		}
	}`
	client := &http.Client{}
	httpReq, _ := http.NewRequest("GET", "http://localhost:9200/chatlog/_search", bytes.NewBufferString(query))
	httpReq.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(httpReq)
	if err != nil {
		log.Print(err)
	}

	respData, err := ioutil.ReadAll(resp.Body)
	w.Write(respData)
}

// GetUserKeywordsHTTP is
func GetUserKeywordsHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	query := `
	{
		"size": 0, 
		"query": {
		  "bool": {
			 "must": [
			   {
				 "term": {
				   "author" : "` + vars["author"] + `"
				 }
			   }
			 ] 
		  }
		},
		"aggs": {
		  "keywords": {
			"terms": {
			  "field": "contents",
			  "order": {
				"_count": "desc"
			  }, 
			  "size": 100
			}
		  }
		}
	  }`
	client := &http.Client{}
	httpReq, _ := http.NewRequest("GET", "http://localhost:9200/chatlog/_search", bytes.NewBufferString(query))
	httpReq.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(httpReq)
	if err != nil {
		log.Print(err)
	}

	respData, err := ioutil.ReadAll(resp.Body)
	w.Write(respData)
}
