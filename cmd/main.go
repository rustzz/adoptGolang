package main

import (
	"adoptGolang/internal/database"
	"adoptGolang/internal/engine"
	"adoptGolang/internal/engine/adopt/handler"
	"adoptGolang/internal/engine/system"
	"adoptGolang/internal/engine/system/sender"
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
	database.Migrate()

	systemHandler := &system.Handler{
		Group: group[0],
		Client: client,
		Sender: &sender.Sender{Client: client},
	}

	controller := &engine.Controller{
		AdoptHandler: &handler.AdoptHandler{ Handler: systemHandler },
	}
	lp.MessageNew(func(_ context.Context, obj events.MessageNewObject) {
		go controller.Handle(obj)
	})

	log.Println("Starting...")
	if err = lp.Run(); err != nil {
		log.Fatal("[FATAL]: ", err)
		return
	}
}
