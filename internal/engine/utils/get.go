package utils

import (
	messageHandler "adoptGolang/internal/engine/handler/message"
	"bytes"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/object"
	"image"
	"image/png"
	"regexp"
	"strings"
)

var (
	mentionPattern = `^\[\w+\d+\|.*\w+\].*$`
)

func GetTexts(message string, countTexts int) (out []string) {
	messageLines := strings.Split(message, "\n")

	fromIndex := 0	/*
					от какого индекса получать текста команды
					ниже: если имеется префикс к боту в виде [club123456|name],
					то отдавать от первого индекса, где находится имя команды
					 */
	matched, _ := regexp.Match(mentionPattern, []byte(strings.ToLower(strings.Split(message, " ")[0])))
	if matched || messageHandler.IsBotNamePrefix(message) {
		fromIndex = 1
	}

	// получить текст с первой строки от имени команды
	_tmp := ""
	for index, word := range strings.Split(messageLines[0], " ") {
		if index > fromIndex {
			_tmp += fmt.Sprintf("%s ", word)
		}
	}

	// добавить строки
	if len(_tmp) > 0 { out = append(out, _tmp[0:len(_tmp)-1]) }
	for _, elem := range messageLines[1:] {
		out = append(out, elem)
	}
	// не достающие строки заполнить пустотой воизбежание нехватки индексов
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
func GetImages(message object.MessagesMessage, countImages int) (srcImageBuffers []*bytes.Buffer, err error) {
	// in message
	if err = messageHandler.ImageExist(message.Attachments); err == nil {
		for _, attachment := range message.Attachments {
			lenOfAvailablePhotos := len(attachment.Photo.Sizes) - 1 /*
				Индекс последней ссылки
				(с макс. разрешением)
			*/
			srcImageBuffer, err := LoadSrcImageBufferFromURL(attachment.Photo.Sizes[lenOfAvailablePhotos].URL)
			if err != nil { return nil, err }
			srcImageBuffers = append(srcImageBuffers, srcImageBuffer)
		}
		return srcImageBuffers, nil
	}
	// in forward
	if err = messageHandler.FwdMessageExist(message); err == nil {
		if err = messageHandler.ImageExist(message.FwdMessages[0].Attachments); err == nil {
			for _, attachment := range message.FwdMessages[0].Attachments {
				lenOfAvailablePhotos := len(attachment.Photo.Sizes) - 1
				srcImageReader, err := LoadSrcImageBufferFromURL(attachment.Photo.Sizes[lenOfAvailablePhotos].URL)
				if err != nil { return nil, err }
				srcImageBuffers = append(srcImageBuffers, srcImageReader)
			}
			return srcImageBuffers, nil
		} else { return nil, err }
	}
	// in reply
	if err = messageHandler.ReplyMessageExist(message); err == nil {
		if err = messageHandler.ImageExist(message.ReplyMessage.Attachments); err == nil {
			for _, attachment := range message.ReplyMessage.Attachments {
				lenOfAvailablePhotos := len(attachment.Photo.Sizes) - 1
				srcImageBuffer, err := LoadSrcImageBufferFromURL(attachment.Photo.Sizes[lenOfAvailablePhotos].URL)
				if err != nil {
					return nil, err
				}
				srcImageBuffers = append(srcImageBuffers, srcImageBuffer)
			}
		} else { return nil, err }
	}

	// заглушка | необходимо для избежания ошибки с nil-pointer
	_tmp := image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: 500, Y: 500}})
	imageBuffer := new(bytes.Buffer)
	if err = png.Encode(imageBuffer, _tmp); err != nil { return }
	for i := len(srcImageBuffers); i < countImages; i++ {
		srcImageBuffers = append(srcImageBuffers, imageBuffer)
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
	matched, _ := regexp.Match(mentionPattern, []byte(strings.ToLower(strings.Split(message, " ")[0])))
	if matched || messageHandler.IsBotNamePrefix(message) {
		if len(strings.Split(message, " ")) < 2 { return "" }
		return strings.ToLower(strings.Split(message, " ")[1])
	}
	return strings.ToLower(strings.Split(message, " ")[0])
}
