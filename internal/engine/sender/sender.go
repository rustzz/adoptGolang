package sender

import (
	"bytes"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
)

type Sender struct {
	Client		*api.VK
	Uploader	*Uploader
}

func (sender *Sender) SendWithImage(peerID int, message string, imageReader *bytes.Reader) (err error) {
	resp, err := sender.uploadImage(peerID, imageReader)
	if err != nil {
		return
	}
	attachment := fmt.Sprintf("photo%d_%d", resp[0].OwnerID, resp[0].ID)
	parameters := params.NewMessagesSendBuilder()
	parameters.PeerID(peerID)
	parameters.Message(message)
	parameters.RandomID(0)
	parameters.Attachment(attachment)

	_, err = sender.Client.MessagesSend(parameters.Params)
	if err != nil {
		return
	}
	return
}

func (sender *Sender) Send(peerID int, message string) (err error) {
	parameters := params.NewMessagesSendBuilder()
	parameters.PeerID(peerID)
	parameters.Message(message)
	parameters.RandomID(0)

	_, err = sender.Client.MessagesSend(parameters.Params)
	if err != nil {
		return
	}
	return
}