package checkers

import (
	"adoptGolang"
	"adoptGolang/internal/engine/adopt/errors"
	"regexp"
	"strings"

	"github.com/SevereCloud/vksdk/v2/object"
)

func IsBotNamePrefix(message string) (matched bool) {
	_tmp := strings.ToLower(strings.Split(message, " ")[0])
	matched, _ = regexp.Match(adoptGolang.Prefix, []byte(_tmp))
	return
}

func FwdMessageExist(message object.MessagesMessage) error {
	if len(message.FwdMessages) > 0 {
		return nil
	}
	return &errors.FwdReplyMessageNotFound{}
}

func ReplyMessageExist(message object.MessagesMessage) error {
	if message.ReplyMessage != nil {
		return nil
	}
	return &errors.FwdReplyMessageNotFound{}
}
