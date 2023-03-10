package main

import (
	"awesomeProject/clients/telegram"
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	tgClient := telegram.New(tgBotHost, mustToken())

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
