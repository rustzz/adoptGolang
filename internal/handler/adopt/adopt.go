package adopt

import (
	"adoptGolang/internal/engine/controller"
	"adoptGolang/internal/engine/system/utils/logger"
	"github.com/SevereCloud/vksdk/v2/events"
)

type HandlerAdopt struct {
	*controller.Client
	*logger.Logger
}

func (handler *HandlerAdopt) HandleTestEcho(obj events.MessageNewObject) (err error) {
	// вызов можно сократить, но я указываю явный путь до метода на всякий случай
	handler.Controller.CustomAPI.Send(
		obj.Message.PeerID, obj.Message.Text, [10][]byte{}, /* пустой массив изображений */
	)
	return
}
