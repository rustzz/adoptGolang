package converter

import "bytes"

func ByteToReader(data []byte) *bytes.Reader {
	return bytes.NewReader(data)
}
