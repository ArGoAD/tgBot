package main

import (
	"flag"
	"log"
)

func main() {

	t := mustToken()

}

func mustToken() string {

	token := flag.String(
		"bot-token",
		"",
		"token for access to tgBot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token

}
