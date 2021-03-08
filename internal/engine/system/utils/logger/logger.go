package logger

import (
	"adoptGolang/internal/engine/controller"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/events"
	"io"
	"os"
)

var (
	logLevels = [4]string{"INFO", "WARNING", "ERROR", "FATAL"}
	logPatterns = [1]string{
		"[%s][%s][%s][CHAT: %d][USER: %d]: %s\n",
	}
)

type Logger struct {
	*controller.Controller
}

func NewLogger(_controller *controller.Controller) *Logger {
	return &Logger{ Controller: _controller }
}

func (logger *Logger) DoLog(
	module func(obj events.MessageNewObject) error,
	logLevel int, moduleName string,
) func(obj events.MessageNewObject) {
	return func(obj events.MessageNewObject) {
		logString := fmt.Sprintf(
			logPatterns[0],
			logLevels[logLevel], "Начало", moduleName,
			obj.Message.PeerID, obj.Message.FromID,
			obj.Message.Text)
		io.WriteString(os.Stdout, logString)
		if err := module(obj); err != nil {
			logString = fmt.Sprintf(
				logPatterns[0],
				logLevels[2], "Процесс", moduleName,
				obj.Message.PeerID, obj.Message.FromID,
				err.Error())
			io.WriteString(os.Stderr, logString)
		}
		logString = fmt.Sprintf(
			logPatterns[0],
			logLevels[logLevel], "Конец", moduleName,
			obj.Message.PeerID, obj.Message.FromID,
			obj.Message.Text)
		io.WriteString(os.Stdout, logString)
	}
}

func (logger *Logger) Log(message string) {
	io.WriteString(os.Stdout, message + "\n")
}
