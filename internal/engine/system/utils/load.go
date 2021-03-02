package utils

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
)

func LoadSrcImageBuffer(path string) (imageBuffer *bytes.Buffer, err error) {
	imageFile, err := os.Open(path)
	if err != nil { return }
	imageBuffer = new(bytes.Buffer)
	imageReader := bufio.NewReader(imageFile)
	if _, err = imageBuffer.ReadFrom(imageReader); err != nil { return }
	return
}

func LoadSrcImageBufferFromURL(url string) (imageBuffer *bytes.Buffer, err error) {
	resp, err := http.Get(url)
	if err != nil { return }
	defer resp.Body.Close()
	imageBytes, err := ioutil.ReadAll(resp.Body)
	imageBuffer = bytes.NewBuffer(imageBytes)
	return
}