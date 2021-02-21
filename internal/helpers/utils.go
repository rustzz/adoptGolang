package helpers

import (
	"bytes"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/object"
	"github.com/rustzz/demotivator"
	"image"
	"image/png"
	"regexp"
	"strings"
)

func GetTexts(message string, countTexts int) (out []string) {
	splitedMessage := strings.Split(message, "\n")
	var _tmp string

	fromIndex := 0
	if matched, _ := regexp.Match(`^\[\s+\d+|@\s+\]$`, []byte(strings.Split(message, " ")[0])); matched {
		fromIndex = 1
	}

	for index, word := range strings.Split(splitedMessage[0], " ") {
		if index > fromIndex {
			_tmp += fmt.Sprintf("%s ", word)
		}
	}
	out = append(out, _tmp)
	for _, elem := range splitedMessage[1:] {
		out = append(out, elem)
	}
	if len(out) != countTexts {
		for i := len(out); i < countTexts; i++ { out = append(out, "") }
	}
	return
}

func GetImages(message object.MessagesMessage, countImages int) (srcImageReaders []*bytes.Reader, err error) {
	// in message
	if err = ImageExist(message.Attachments); err == nil {
		for _, attachment := range message.Attachments {
			lenOfAvailablePhotos := len(attachment.Photo.Sizes) - 1
			srcImageReader, err := demotivator.LoadSrcImageFromURL(
				attachment.Photo.Sizes[lenOfAvailablePhotos].URL,
			)
			if err != nil { return nil, err }
			srcImageReaders = append(srcImageReaders, srcImageReader)
		}
		return srcImageReaders, nil
	}
	// in forward
	if err = FwdMessageExist(message); err == nil {
		if err = ImageExist(message.FwdMessages[0].Attachments); err == nil {
			for _, attachment := range message.FwdMessages[0].Attachments {
				lenOfAvailablePhotos := len(attachment.Photo.Sizes) - 1
				srcImageReader, err := demotivator.LoadSrcImageFromURL(
					attachment.Photo.Sizes[lenOfAvailablePhotos].URL,
				)
				if err != nil { return nil, err }
				srcImageReaders = append(srcImageReaders, srcImageReader)
			}
			return srcImageReaders, nil
		} else { return nil, err }
	}
	// in reply
	if err = ReplyMessageExist(message); err == nil {
		if err = ImageExist(message.ReplyMessage.Attachments); err == nil {
			for _, attachment := range message.ReplyMessage.Attachments {
				lenOfAvailablePhotos := len(attachment.Photo.Sizes) - 1
				srcImageReader, err := demotivator.LoadSrcImageFromURL(
					attachment.Photo.Sizes[lenOfAvailablePhotos].URL,
				)
				if err != nil {
					return nil, err
				}
				srcImageReaders = append(srcImageReaders, srcImageReader)
			}
		} else { return nil, err }
	}

	// заглушка
	_tmp := image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: 100, Y: 100}})
	imageBytes := &bytes.Buffer{}
	png.Encode(imageBytes, _tmp)
	srcImageReader := bytes.NewReader(imageBytes.Bytes())
	if len(srcImageReaders) != countImages {
		for i := len(srcImageReaders); i < countImages; i++ {
			srcImageReaders = append(srcImageReaders, srcImageReader)
		}
	}
	return
}

func GetCommand(message string) string {
	if matched, _ := regexp.Match(`^\[\s+\d+|@\s+\]$`, []byte(strings.Split(message, " ")[0])); matched {
		return strings.Split(message, " ")[1]
	}
	return strings.Split(message, " ")[0]
}
