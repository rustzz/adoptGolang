package engine

import (
	"adoptGolang/internal/engine/adopt/handler"
	messageHandler "adoptGolang/internal/engine/adopt/handler/text"
	"adoptGolang/internal/engine/system/logger"
	"adoptGolang/internal/engine/system/utils"
	"fmt"
	"log"

	"github.com/SevereCloud/vksdk/v2/events"
)

// основной контроллер обработчиков
// можно добавлять дополнения для бота
type Controller struct {
	AdoptHandler  *handler.AdoptHandler
}

func (controller *Controller) Handle(obj events.MessageNewObject) {
	log.Println(
		fmt.Sprintf("[INFO][OUT][CHAT: %d][USER: %d]: %s",
			obj.Message.PeerID, obj.Message.FromID, obj.Message.Text,
		))

	command := utils.GetCommand(obj.Message.Text)
	if messageHandler.IsDem(command) {
		wLog := logger.WrapLog(controller.AdoptHandler.HandleDem, 0, "Демотиватор")
		wLog(obj)
		return // решил добавить, воизбежание double-text
	}
	if messageHandler.IsTBD(command) {
		wLog := logger.WrapLog(controller.AdoptHandler.HandleTBD, 0, "TBD")
		wLog(obj)
		return
	}
	if messageHandler.IsLiquidRescale(command) {
		wLog := logger.WrapLog(controller.AdoptHandler.HandleLiquidRescale, 0, "LiquidRescale")
		wLog(obj)
		return
	}
	return
}
