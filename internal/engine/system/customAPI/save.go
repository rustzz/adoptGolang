package customAPI

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
)


// saveImage : загрузка изображения для отправки в чат
// Todo : В будущем открыть для использования во всех случаях надобности
func (cAPI *CustomAPI) saveImage(peerID int, imgBytes []byte) (
	resp api.PhotosSaveMessagesPhotoResponse, err error,
) {
	server, photo, hash, err := func () (server int, photo, hash string, err error) {
		parameters := params.NewPhotosGetMessagesUploadServerBuilder()
		parameters.PeerID(peerID)
		resp, err := cAPI.API.PhotosGetMessagesUploadServer(parameters.Params)
		if err != nil { return }
		server, photo, hash, err = cAPI.uploadImage(resp.UploadURL, imgBytes)
		if err != nil { return }
		return
	}()
	if err != nil { return }
	parameters := params.NewPhotosSaveMessagesPhotoBuilder()
	parameters.Server(server)
	parameters.Photo(photo)
	parameters.Hash(hash)

	resp, err = cAPI.API.PhotosSaveMessagesPhoto(parameters.Params)
	if err != nil { return }
	return
}