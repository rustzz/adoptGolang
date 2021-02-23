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
	log.Println(fmt.Sprintf(
			"[INFO][OUT][CHAT: %d][USER: %d]: %s",
			obj.Message.PeerID, obj.Message.FromID, obj.Message.Text))

	command := engine.GetCommand(obj.Message.Text)
	if messageHandler.IsDem(command) {
		log.Println(fmt.Sprintf(
			"[INFO][Начало][Демотиватор][CHAT: %d][USER: %d]: %s",
			obj.Message.PeerID, obj.Message.FromID, obj.Message.Text))
		handler.HandleDem(obj)
		log.Println(fmt.Sprintf(
			"[INFO][Конец][Демотиватор][CHAT: %d][USER: %d]: %s",
			obj.Message.PeerID, obj.Message.FromID, obj.Message.Text))
		return // решил добавить, воизбежание double-message
	}
	if messageHandler.IsTBD(command) {
		log.Println(fmt.Sprintf(
			"[INFO][Начало][TBD][CHAT: %d][USER: %d]: %s",
			obj.Message.PeerID, obj.Message.FromID, obj.Message.Text))
		handler.HandleTBD(obj)
		log.Println(fmt.Sprintf(
			"[INFO][Начало][TBD][CHAT: %d][USER: %d]: %s",
			obj.Message.PeerID, obj.Message.FromID, obj.Message.Text))
		return
	}
	if messageHandler.IsLiquidRescale(command) {
		log.Println(fmt.Sprintf(
			"[INFO][Начало][LiquidRescale][CHAT: %d][USER: %d]: %s",
			obj.Message.PeerID, obj.Message.FromID, obj.Message.Text))
		handler.HandleLiquidRescale(obj)
		log.Println(fmt.Sprintf(
			"[INFO][Начало][LiquidRescale][CHAT: %d][USER: %d]: %s",
			obj.Message.PeerID, obj.Message.FromID, obj.Message.Text))
		return
	}
	return
}
