package top

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

// GetDateRangeHTTP is
func GetDateRangeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	// from, _ := r.URL.Query()["from"]
	// size, _ := r.URL.Query()["size"]

	query := `
	{
		"size" : 0,
		"query": {
		  "match_all": {} 
		},
		"aggs": {
		  "maxToDate": {
			"max": {
			  "field": "date"
			}
		  },
		  "minFromDate" : {
			"min": {
			  "field": "date"
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

// GetChatLogHTTP is
func GetChatLogHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	from, _ := r.URL.Query()["from"]
	size, _ := r.URL.Query()["size"]
	fromDate, _ := r.URL.Query()["fromDate"]
	toDate, _ := r.URL.Query()["toDate"]

	query := `
	{
		"size" : ` + size[0] + `,
		"from": ` + from[0] + `, 
		"query": {
			"bool": {
				"must": [
					{
						"range": {
							"date": {
								"gte": "` + fromDate[0] + `",
								"lte": "` + toDate[0] + `",
								"format": "yyyy-MM-dd"
							}
						}
					}
				]
			} 
		},
		"sort": [
			{
				"date": {
					"order": "asc"
				}
			},
			{
				"_id" : {
					"order": "asc"
				}
			}
		]
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

// GetDateChartKeywordHTTP is
func GetDateChartKeywordHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fromDate, _ := r.URL.Query()["fromDate"]
	toDate, _ := r.URL.Query()["toDate"]
	// from, _ := r.URL.Query()["from"]
	// size, _ := r.URL.Query()["size"]
	query := `
	{
		"size": 0, 
		"query": {
			"bool": {
				"must": [
					{
						"range": {
							"date": {
								"gte": "` + fromDate[0] + `",
								"lte": "` + toDate[0] + `",
								"format": "yyyy-MM-dd"
							}
						}
					}
				]
			} 
		},
		"aggs": {
			"keyword": {
				"terms": {
					"field": "contents",
					"size" : 100
				}
			}
		}
	}`
	log.Print(query)
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

// GetDateChartHTTP is
func GetDateChartHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fromDate, _ := r.URL.Query()["fromDate"]
	toDate, _ := r.URL.Query()["toDate"]
	interval, _ := r.URL.Query()["interval"]
	// from, _ := r.URL.Query()["from"]
	// size, _ := r.URL.Query()["size"]
	format := "yy/MM/dd"
	if interval[0] == "month" {
		format = "yy/MM"
	}
	query := `
	{
		"size": 0, 
		"query": {
			"bool": {
				"must": [
					{
						"range": {
							"date": {
								"gte": "` + fromDate[0] + `",
								"lte": "` + toDate[0] + `",
								"format": "yyyy-MM-dd"
							}
						}
					}
				]
			} 
		},
		"aggs": {
			"daily": {
				"date_histogram": {
					"field": "date",
					"interval": "` + interval[0] + `",
					"format": "` + format + `"
				},
				"aggs": {
					"author": {
						"terms": {
							"field": "author"
						}
					}
				}
			}
		}
	}`
	log.Print(query)
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
