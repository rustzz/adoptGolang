package handler

import "C"
import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/rustzz/adopt/internal/helpers"
	"github.com/rustzz/adopt/internal/sender"
	"log"
)

type Handler struct {
	Client	*api.VK
	Sender	*sender.Sender
}

func (handler *Handler) Handle(obj *events.MessageNewObject) {
	log.Println("[INFO]: ", fmt.Sprintf("%d: %s", obj.Message.PeerID, obj.Message.Text))

	command := helpers.GetCommand(obj.Message.Text)
	if helpers.IsDem(command) {
		handler.HandleDem(obj)
	}
	if helpers.IsTBD(command) {
		handler.HandleTBD(obj)
	}
	if helpers.IsСumCas(command) {
		handler.HandleTBD(obj)
	}
	return
}
