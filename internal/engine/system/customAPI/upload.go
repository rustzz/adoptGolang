package customAPI

import (
	"adoptGolang/internal/engine/system/utils/converter"
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
)

type uploadResponse struct {
	Server int    `json:"server"`
	Photo  string `json:"photo"`
	Hash   string `json:"hash"`
}

// uploadImage : отправка изображения на сервер
func (cAPI *CustomAPI) uploadImage(url string, img []byte) (
	server int, photo, hash string, err error,
) {
	multipartBuffer := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(multipartBuffer)

	fileWriter, err := multipartWriter.CreateFormFile("photo", "photo.jpg")
	if err != nil { return }
	imgReader := converter.ByteToReader(img)
	imgReader.Seek(0, 0)
	if _, err = io.Copy(fileWriter, imgReader); err != nil { return }
	if err = multipartWriter.Close(); err != nil { return }

	req, err := http.NewRequest("POST", url, multipartBuffer)
	if err != nil { return }
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil { return }

	decodedData := json.NewDecoder(resp.Body)
	response := &uploadResponse{}
	if err = decodedData.Decode(response); err != nil { return }
	defer resp.Body.Close()

	server = response.Server
	photo = response.Photo
	hash = response.Hash
	return
}
