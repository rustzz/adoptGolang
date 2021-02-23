package main

import (
	"adoptGolang/internal/engine/handler"
	"adoptGolang/internal/engine/sender"
	"context"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"log"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	token := os.Getenv("VK_GROUP_TOKEN")
	client := api.NewVK(token)

	// настройка лонгполла для бота
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
