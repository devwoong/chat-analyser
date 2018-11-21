package main

import (
	"chat-analyser/elastic-bulk/parser"
)

func main() {
	parser.Parse("./kakao.txt")
}
