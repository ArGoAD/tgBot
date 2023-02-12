package telegram

import "awesomeProject/clients/telegram"

//This struct init interface Processor

type Processor struct {
	tg     *telegram.Client
	offset int
	//storage

}

//To create a processor

func New(client *telegram.Client) {}
