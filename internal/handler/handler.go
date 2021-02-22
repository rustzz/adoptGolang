package handler

import "C"
import (
	"adoptGolang/internal/helpers"
	"adoptGolang/internal/sender"
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

	command := helpers.GetCommand(obj.Message.Text)
	if helpers.IsDem(command) {
		handler.HandleDem(obj)
		return // решил добавить, воизбежание double-message
	}
	if helpers.IsTBD(command) {
		handler.HandleTBD(obj)
		return
	}
	if helpers.IsLiquidRescale(command) {
		handler.HandleLiquidRescale(obj)
		return
	}
	return
}
