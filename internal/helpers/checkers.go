package helpers

import (
	"adoptGolang"
	"adoptGolang/internal/errors"
	"github.com/SevereCloud/vksdk/v2/object"
	"regexp"
	"strings"
)

func IsBotNamePrefix(message string) (matched bool) {
	_tmp := strings.ToLower(strings.Split(message, " ")[0])
	matched, _ = regexp.Match(adoptGolang.Prefix, []byte(_tmp))
	return
}

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
