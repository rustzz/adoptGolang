package customAPI

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
)

type CustomAPI struct {
	GroupID	int
	API		*api.VK
}

func (cAPI *CustomAPI) Send(peerID int, message string, images [10][]byte) (err error) {
	parameters := params.NewMessagesSendBuilder()
	parameters.RandomID(0)
	parameters.PeerID(peerID)
	parameters.Message(message)
	var attachments string
	for _, image := range images {
		if image != nil {
			resp, err := cAPI.saveImage(peerID, image)
			if err != nil { return err }
			attachments += fmt.Sprintf("photo%d_%d,", resp[0].OwnerID, resp[0].ID)
		}
	}
	{
		parameters := params.NewMessagesSendBuilder()
		parameters.PeerID(peerID)
		parameters.Message(message)
		parameters.RandomID(0)
		parameters.Attachment(attachments)

		_, err = cAPI.API.MessagesSend(parameters.Params)
		if err != nil { return }
	}
	return
}
