package handler

import "C"
import (
	"adoptGolang/internal/engine"
	messageHandler "adoptGolang/internal/engine/handler/message"
	"adoptGolang/internal/engine/sender"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
	"log"
)

type Handler struct {
	Client	*api.VK
	Sender	*sender.Sender /*
							чтобы не передавать каждой бляди как аргумент
							и не мазолило глаз
							*/
}

func (handler *Handler) Handle(obj *events.MessageNewObject) {
	log.Println("[INFO]: ", fmt.Sprintf("%d: %s", obj.Message.PeerID, obj.Message.Text))

	command := engine.GetCommand(obj.Message.Text)
	if messageHandler.IsDem(command) {
		handler.HandleDem(obj)
		return // решил добавить, воизбежание double-message
	}
	if messageHandler.IsTBD(command) {
		handler.HandleTBD(obj)
		return
	}
	if messageHandler.IsLiquidRescale(command) {
		handler.HandleLiquidRescale(obj)
		return
	}
	return
}
