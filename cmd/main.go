package main

import (
	"context"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"github.com/rustzz/adopt/internal/handler"
	"github.com/rustzz/adopt/internal/sender"
	"log"
	"os"
)

func main() {
	token := os.Getenv("VK_GROUP_TOKEN")
	client := api.NewVK(token)

	group, err := client.GroupsGetByID(nil)
	if err != nil {
		log.Fatal("[FATAL]: ", err)
		return
	}
	lp, err := longpoll.NewLongPoll(client, group[0].ID)
	if err != nil {
		log.Fatal("[FATAL]: ", err)
		return
	}

	Handler := &handler.Handler{
		Client: client,
		Sender: &sender.Sender{Client: client},
	}
	lp.MessageNew(func(_ context.Context, obj events.MessageNewObject) {
		go Handler.Handle(&obj)
	})

	log.Println("Starting...")
	if err = lp.Run(); err != nil {
		log.Fatal("[FATAL]: ", err)
		return
	}
}
