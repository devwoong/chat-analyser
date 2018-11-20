#/bin/sh

curl -XPOST localhost:9200/chatlog/_bulk -H "Content-Type:application/x-ndjson" --data-binary @"$1"