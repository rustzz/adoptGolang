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

	fromIndex := 0	/*
					от какого индекса получать текста команды
					ниже: если имеется префикс к боту в виде [club123456|name],
					то отдавать от первого индекса, где находится имя команды
					*/
	_tmp = strings.ToLower(strings.Split(message, " ")[0])
	matched, _ := regexp.Match(`^\[\w+\d+\|.*\w+\].*$`, []byte(_tmp))
	if matched || IsBotNamePrefix(message) {
		fromIndex = 1
	}

	_tmp = ""
	for index, word := range strings.Split(splitedMessage[0], " ") {
		if index > fromIndex {
			_tmp += fmt.Sprintf("%s ", word)
		}
	}
	out = append(out, _tmp[0:len(_tmp)-1])
	for _, elem := range splitedMessage[1:] {
		out = append(out, elem)
	}
	if len(out) != countTexts {
		for i := len(out); i < countTexts; i++ { out = append(out, "") }
	}
	return
}

/*
GetImages : Собирает все изображения из сообщения
			- 1: в сообщении
			- 2: в пересланном сообщении (по первому индексу; не со всех пересланных)
			- 3: в отвеченном сообщении
*/
func GetImages(message object.MessagesMessage, countImages int) (srcImageReaders []*bytes.Reader, err error) {
	// in message
	if err = ImageExist(message.Attachments); err == nil {
		for _, attachment := range message.Attachments {
			lenOfAvailablePhotos := len(attachment.Photo.Sizes) - 1 /*
																	Индекс последней ссылки
																	(с макс. разрешением)
																	 */
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

	// заглушка | необходимо для избежания ошибки с nil-pointer
	_tmp := image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: 100, Y: 100}})
	imageBytes := &bytes.Buffer{}
	png.Encode(imageBytes, _tmp)
	srcImageReader := bytes.NewReader(imageBytes.Bytes())
	for i := len(srcImageReaders); i < countImages; i++ {
		srcImageReaders = append(srcImageReaders, srcImageReader)
		if i > countImages { return }	/*
								 закончить, если при сборе изображений
								 выше их количество было больше запрошенного
								 */
	}
	return
}

/*
GetCommand : Отдает имя команды в зависимости как вызывался бот
			[club123456|name], -  с запятой для моб. клиента
								  без запятой для веб клиента
 */
func GetCommand(message string) string {
	_tmp := strings.Split(message, " ")[0]
	matched, _ := regexp.Match(`^\[\w+\d+\|.*\w+\].*$`, []byte(strings.ToLower(_tmp)))
	if matched || IsBotNamePrefix(message) {
		if len(strings.Split(message, " ")) < 2 { return "" }
		return strings.ToLower(strings.Split(message, " ")[1])
	}
	return strings.ToLower(strings.Split(message, " ")[0])
}
