package image

import (
	"io/ioutil"
	"net/http"
	"time"
)

func LoadSrcImageBytesFromURL(url string) (imgBytes []byte, err error) {
	client := &http.Client{ Timeout: 30 * time.Second }
	resp, err := client.Get(url)
	if err != nil { return }
	defer resp.Body.Close()
	imgBytes, err = ioutil.ReadAll(resp.Body)
	return
}
