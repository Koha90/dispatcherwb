// Package main general for bot
package main

import (
	"context"
	"flag"
	"log"

	tgClient "github.com/koha90/telegram-bot-api/clients/telegram"
	event_consumer "github.com/koha90/telegram-bot-api/consumer/event-consumer"
	"github.com/koha90/telegram-bot-api/events/telegram"
	"github.com/koha90/telegram-bot-api/storage/sqlite"
)

const (
	host           = "api.telegram.org"
	storagePath    = "storagefile"
	storageSqlPath = "data/sqlite/storage.db"
	batchSize      = 100
)

func main() {
	// s := files.New(storagePath)
	s, err := sqlite.New(storageSqlPath)
	if err != nil {
		log.Fatalf("can't connect to storage: ", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("can't init storage: ", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(host, mustToken()),
		s,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal(err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
