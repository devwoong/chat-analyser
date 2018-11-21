#/bin/sh
curl -XPUT -H "Content-Type: application/json; charset=utf-8" localhost:9200/chatlog -d '
{
    "settings" : {
        "number_of_shards" : 5,
        "number_of_replicas" : 1,
        "analysis": {
            "analyzer": {
                "arirang": { 
                    "type": "arirang_analyzer",
                    "tokenizer": "arirang_tokenizer",
                    "filter": [
                        "lowercase", "trim", "arirang_filter"]
                },
                "arirang_query" : {
                    "type" : "arirang_analyzer",
                    "queryMode" : true
                }
            }
        }
    },
    "mappings" : {
        "log" : {
            "properties" : {
                "date" : {
                    "type" : "date",
                    "format"  : "yyyy-MM-dd HH:mm",
                    "index" : "true"
                },
                "author" : {
                    "type" : "text",
                    "index" : "true",
                    "fielddata" : "true"
                },
                "contents" : {
                    "type" : "text",
                    "index" : "true",
                    "analyzer" : "arirang",
                    "search_analyzer" : "arirang_query",
                    "fielddata" : "true"
                }
            }
        }
    } 
}'