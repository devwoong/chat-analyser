#/bin/sh

 for file in ./*.json; 
 do 
    echo `basename $file`;
    curl -XPOST localhost:9200/chatlog/_bulk -H "Content-Type:application/x-ndjson" --data-binary @`basename $file`
done
# curl -XPOST localhost:9200/chatlog/_bulk -H "Content-Type:application/x-ndjson" --data-binary @"$1"