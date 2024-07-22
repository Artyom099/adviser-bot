package main

import (
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	token := mustToken()

	tgClient := telegram.New(token, tgBotHost)

	// отправляет запрос в тг, чтоб получать новые события
	// fetcher = fetcher.New(tgClient)

	// отправляет нам новые сообщения после обработки
	// processor = processor.New(tgClient)

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String("bot-token", "", "telegram bot access token")
	flag.Parse()
	if *token == "" {
		//panic("missing bot access token")
		log.Fatal("missing bot access token")
	}
}
