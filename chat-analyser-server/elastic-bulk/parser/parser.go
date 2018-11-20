package parser

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

// Parse is convert Log file to elastic json file
func Parse(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writeFile, err := os.Create("./bulk.json")
	if err != nil {
		log.Fatal(err)
	}
	defer writeFile.Close()

	writer := bufio.NewWriter(writeFile)

	scanner := bufio.NewScanner(file)

	count := 1
	fileNum := 1
	for scanner.Scan() {
		contents := new(Kakao)
		err := contents.parse(scanner.Text())
		if err == nil {
			writer.WriteString("{\"index\":{\"_type\":\"log\", \"_id\":" + strconv.Itoa(count) + "}}\n")

			jsonBytes, err := json.Marshal(contents)

			if err != nil {
				log.Fatal(err)
				panic(err)
			}
			jsonString := string(jsonBytes)
			log.Println(jsonString)
			writer.WriteString(jsonString + "\n")
			writer.Flush()
			count++
			if count >= 2000*fileNum {
				writer.Flush()
				writeFile.Close()
				writeFile, _ = os.Create("./bulk" + strconv.Itoa(fileNum) + ".json")
				writer = bufio.NewWriter(writeFile)
				fileNum++
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
