package system

import (
	"adoptGolang/internal/engine/system/sender"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/object"
)

// определяется архитектурная структура обработчика,
// которая может быть наследована другими дополнениями к opensource боту
type Handler struct {
	Group		object.GroupsGroup
	Client		*api.VK
	Sender		*sender.Sender /*
								чтобы не передавать каждой бляди как аргумент
								и не мазолило глаз
								*/
}

