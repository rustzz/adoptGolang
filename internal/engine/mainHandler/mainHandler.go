package mainHandler

import (
	"adoptGolang/internal/engine/controller"
	"adoptGolang/internal/engine/system/middleware"
	"adoptGolang/internal/engine/system/utils/logger"
	"adoptGolang/internal/engine/system/utils/message/text"
	"adoptGolang/internal/handler/adopt"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/object"
)

type MainHandler struct {
	*controller.Controller
	*logger.Logger
	*middleware.Middleware

	*adopt.HandlerAdopt
}

func NewMainHandler (
	client *api.VK, group object.GroupsGroup,
) *MainHandler {
	_controller := controller.NewController(client, group)
	return &MainHandler{
		Controller: _controller,
		Logger: logger.NewLogger(_controller),
		Middleware: middleware.NewMiddleware(_controller),
	}
}

func (mh *MainHandler) Handle(obj events.MessageNewObject) {
	command := text.GetCommand(obj.Message.Text)

	if adopt.IsTestEcho(command) {
		mh.Logger.DoLog(mh.HandlerAdopt.HandleTestEcho, 0, "TestEcho")(obj)
		return
	}
}
