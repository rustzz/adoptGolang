package checker

import (
	"adoptGolang/internal/engine/system/utils/message/checker/errors"

	"github.com/SevereCloud/vksdk/v2/object"
)

func FwdMessageExist(message object.MessagesMessage) error {
	if len(message.FwdMessages) > 0 { return nil }
	return &errors.FwdMessageNotFound{}
}

func RepliedMessageExist(message object.MessagesMessage) error {
	if message.ReplyMessage != nil { return nil }
	return &errors.RepliedMessageNotFound{}
}
