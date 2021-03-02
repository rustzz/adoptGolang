package sender

import (
	"bytes"
	"encoding/json"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"io"
	"mime/multipart"
	"net/http"
)

type UploadResponse struct {
	Server int    `json:"server"`
	Photo  string `json:"photo"`
	Hash   string `json:"hash"`
}

// sendImage : отправка изображения на сервер
func (sender *Sender) sendImage(url string, img io.Reader) (server int, photo, hash string, err error) {
	multipartBuffer := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(multipartBuffer)

	fileWriter, err := multipartWriter.CreateFormFile("photo", "photo.jpg")
	if err != nil { return }
	if _, err = io.Copy(fileWriter, img); err != nil { return }
	if err = multipartWriter.Close(); err != nil { return }

	req, err := http.NewRequest("POST", url, multipartBuffer)
	if err != nil { return }
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil { return }

	decodedData := json.NewDecoder(resp.Body)
	uploadResponse := &UploadResponse{}
	if err = decodedData.Decode(uploadResponse); err != nil { return }
	defer resp.Body.Close()

	server = uploadResponse.Server
	photo = uploadResponse.Photo
	hash = uploadResponse.Hash
	return
}

/*
uploadImage : совмещающая функция по загрузке изображения
	- получает сервер
	- загружает
	- сохраняет
 */
func (sender *Sender) uploadImage(
	peerID int, imageReader *bytes.Reader,
) (resp api.PhotosSaveMessagesPhotoResponse, err error) {
	server, photo, hash, err := func () (server int, photo, hash string, err error) {
		parameters := params.NewPhotosGetMessagesUploadServerBuilder()
		parameters.PeerID(peerID)
		resp, err := sender.Client.PhotosGetMessagesUploadServer(parameters.Params)
		if err != nil { return }
		server, photo, hash, err = sender.sendImage(resp.UploadURL, imageReader)
		if err != nil { return }
		return
	}()
	if err != nil { return }
	parameters := params.NewPhotosSaveMessagesPhotoBuilder()
	parameters.Server(server)
	parameters.Photo(photo)
	parameters.Hash(hash)

	resp, err = sender.Client.PhotosSaveMessagesPhoto(parameters.Params)
	if err != nil { return }
	return
}
