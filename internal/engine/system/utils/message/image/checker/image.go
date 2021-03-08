package checker

import (
	"adoptGolang/internal/engine/system/utils/message/image/checker/errors"
	"github.com/SevereCloud/vksdk/v2/object"
)

func ImageExist(attachments []object.MessagesMessageAttachment) error {
	for _, attachment := range attachments {
		if len(attachment.Photo.Sizes) > 0 { return nil }
	}
	return &errors.ImageNotFound{}
}
