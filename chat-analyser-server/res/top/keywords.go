package top

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

// GetTotalTop100KeywordsHTTP is
func GetTotalTop100KeywordsHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	// from, _ := r.URL.Query()["from"]
	// size, _ := r.URL.Query()["size"]

	query := `
	{
		"size": 0, 
		"query": {
			"match_all": {}
		},
		"aggs": {
			"keywords": {
				"terms": {
					"field": "contents",
					"order": {
						"_count": "desc"
					}, 
					"size": 100
				},
				"aggs": {
					"author": {
						"terms": {
							"field": "author",
							"size": 100
						}
					}
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
