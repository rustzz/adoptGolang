package image

import (
	imageUtils "adoptGolang/internal/engine/system/utils/image"
	messageChecker "adoptGolang/internal/engine/system/utils/message/checker"
	imageChecker "adoptGolang/internal/engine/system/utils/message/image/checker"
	imageCheckerErrors "adoptGolang/internal/engine/system/utils/message/image/checker/errors"
	"bytes"
	"github.com/SevereCloud/vksdk/v2/object"
	"image"
	"image/png"
)

/*
GetImages : Собирает все изображения из сообщения
			- 1: в сообщении
			- 2: в пересланном сообщении (по первому индексу; не со всех пересланных)
			- 3: в отвеченном сообщении
*/
func GetImages(message object.MessagesMessage, reqCountImages int) (
	srcImgBytesSlice [10][]byte, err error,
) {
	// in message
	if err = imageChecker.ImageExist(message.Attachments); err == nil {
		for index, attachment := range message.Attachments {
			lenOfAvailablePhotos := len(attachment.Photo.Sizes) - 1 /*
																	Индекс последней ссылки
																	(с макс. разрешением)
																	*/
			srcImgBytes, err := imageUtils.LoadSrcImageBytesFromURL(
				attachment.Photo.Sizes[lenOfAvailablePhotos].URL)
			if err != nil { return srcImgBytesSlice, err }
			srcImgBytesSlice[index] = srcImgBytes
		}
		return srcImgBytesSlice, nil
	}
	// in forward
	if err = messageChecker.FwdMessageExist(message); err == nil {
		if err = imageChecker.ImageExist(message.FwdMessages[0].Attachments); err == nil {
			for index, attachment := range message.FwdMessages[0].Attachments {
				lenOfAvailablePhotos := len(attachment.Photo.Sizes) - 1
				srcImgBytes, err := imageUtils.LoadSrcImageBytesFromURL(
					attachment.Photo.Sizes[lenOfAvailablePhotos].URL)
				if err != nil { return srcImgBytesSlice, err }
				srcImgBytesSlice[index] = srcImgBytes
			}
			return srcImgBytesSlice, nil
		} else { return srcImgBytesSlice, err }
	}
	// in reply
	if err = messageChecker.RepliedMessageExist(message); err == nil {
		if err = imageChecker.ImageExist(message.ReplyMessage.Attachments); err == nil {
			for index, attachment := range message.ReplyMessage.Attachments {
				lenOfAvailablePhotos := len(attachment.Photo.Sizes) - 1
				srcImgBytes, err := imageUtils.LoadSrcImageBytesFromURL(
					attachment.Photo.Sizes[lenOfAvailablePhotos].URL)
				if err != nil { return srcImgBytesSlice, err }
				srcImgBytesSlice[index] = srcImgBytes
			}
		} else { return srcImgBytesSlice, err }
	}

	// вообще нет изображений от пользователя
	if srcImgBytesSlice[0] == nil { return srcImgBytesSlice, &imageCheckerErrors.ImageNotFound{} }
	// заглушка | необходимо для избежания ошибки с nil-pointer
	_tmp := image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: 500, Y: 500}})
	imgBuffer := new(bytes.Buffer)
	if err = png.Encode(imgBuffer, _tmp); err != nil { return }
	for i := len(srcImgBytesSlice); i < reqCountImages; i++ {
		srcImgBytesSlice[i] = imgBuffer.Bytes()
		if i > reqCountImages { return }	/*
											 закончить, если при сборе изображений
											 выше их количество было больше запрошенного
											*/
	}
	return
}
