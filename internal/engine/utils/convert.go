package utils

import "bytes"

func BufferToReader(imageBuffer *bytes.Buffer) (imageReader *bytes.Reader) {
	imageReader = bytes.NewReader(imageBuffer.Bytes())
	return
}

func BytesToReader(imageBytes []byte) (imageReader *bytes.Reader) {
	imageReader = bytes.NewReader(imageBytes)
	return
}
