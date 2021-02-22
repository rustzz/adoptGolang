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

// postImage : отправка изображения на сервер
func postImage(url string, img io.Reader) (server int, photo, hash string, err error) {
	type UploadResponse struct {
		Server int    `json:"server"`
		Photo  string `json:"photo"`
		Hash   string `json:"hash"`
	}

	var multipartBuffer bytes.Buffer
	multipartWriter := multipart.NewWriter(&multipartBuffer)
	fileWriter, err := multipartWriter.CreateFormFile("photo", "photo.jpg")
	if err != nil {
		return
	}
	if _, err = io.Copy(fileWriter, img); err != nil {
		return
	}
	multipartWriter.Close()

	req, err := http.NewRequest("POST", url, &multipartBuffer)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	uploadResp := UploadResponse{}
	decodedData := json.NewDecoder(resp.Body)
	err = decodedData.Decode(&uploadResp)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	server = uploadResp.Server
	photo = uploadResp.Photo
	hash = uploadResp.Hash
	return
}

/*
uploadImage : совмещающая функция по загрузке изображения
	- получает сервер
	- загружает
	- сохраняет
 */
func (sender *Sender) uploadImage(peerID int, imageReader *bytes.Reader) (resp api.PhotosSaveMessagesPhotoResponse, err error) {
	server, photo, hash, err := func () (server int, photo, hash string, err error) {
		parameters := params.NewPhotosGetMessagesUploadServerBuilder()
		parameters.PeerID(peerID)
		resp, err := sender.Client.PhotosGetMessagesUploadServer(parameters.Params)
		if err != nil {
			return
		}
		server, photo, hash, err = postImage(resp.UploadURL, imageReader)
		if err != nil {
			return
		}
		return
	}()
	if err != nil {
		return
	}
	parameters := params.NewPhotosSaveMessagesPhotoBuilder()
	parameters.Server(server)
	parameters.Photo(photo)
	parameters.Hash(hash)

	resp, err = sender.Client.PhotosSaveMessagesPhoto(parameters.Params)
	if err != nil {
		return
	}
	return
}
