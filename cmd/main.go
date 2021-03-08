package main

import (
	"adoptGolang/internal/engine/mainHandler"
	"adoptGolang/internal/handler/adopt"
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

	// миграция бд
	//database.Migrate()

	handler := mainHandler.NewMainHandler(client, group[0])
	handler.HandlerAdopt = &adopt.HandlerAdopt{
		Client: handler.Controller.Client,
		Logger: handler.Logger,
	}

	lp.MessageNew(func(_ context.Context, obj events.MessageNewObject) {
		go handler.Handle(obj)
	})

	log.Println("Starting...")
	if err = lp.Run(); err != nil {
		log.Fatal("[FATAL]: ", err)
		return
	}
}
