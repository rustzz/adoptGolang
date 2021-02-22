package helpers

import (
	"github.com/SevereCloud/vksdk/v2/object"
	"adoptGolang/internal/errors"
)

func ImageExist(attachments []object.MessagesMessageAttachment) error {
	for _, attachment := range attachments {
		if len(attachment.Photo.Sizes) > 0 { return nil }
	}
	return &errors.ImageNotFound{}
}

func FwdMessageExist(message object.MessagesMessage) error {
	if len(message.FwdMessages) > 0 { return nil }
	return &errors.FwdReplyMessageNotFound{}
}

func ReplyMessageExist(message object.MessagesMessage) error {
	if message.ReplyMessage != nil { return nil }
	return &errors.FwdReplyMessageNotFound{}
}
