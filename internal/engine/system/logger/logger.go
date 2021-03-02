package logger

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/events"
	"log"
)

var (
	logLevels = [4]string{"INFO", "WARNING", "ERROR", "FATAL"}
)

func WrapLog(
	f func(obj events.MessageNewObject),
	logLevel int, module string,
) func(obj events.MessageNewObject) {
	return func(obj events.MessageNewObject) {
		log.Println(fmt.Sprintf(
			"[%s][Начало][%s][CHAT: %d][USER: %d]: %s", logLevels[logLevel], module,
			obj.Message.PeerID, obj.Message.FromID, obj.Message.Text))
		f(obj)
		log.Println(fmt.Sprintf(
			"[%s][Конец][%s][CHAT: %d][USER: %d]: %s", logLevels[logLevel], module,
			obj.Message.PeerID, obj.Message.FromID, obj.Message.Text))
	}
}
