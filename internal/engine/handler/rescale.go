package handler

import (
	"adoptGolang/internal/engine/utils"
	"github.com/SevereCloud/vksdk/v2/events"
	"gopkg.in/gographics/imagick.v3/imagick"
	"image"
	"log"
	"strconv"
)

// Todo : Перенести в внешний модуль (репозиторий)
func DoRescale(handler *Handler, obj *events.MessageNewObject) {
	imagick.Initialize()
	defer imagick.Terminate()

	texts := utils.GetTexts(obj.Message.Text, 1)
	countOfRescales, err := strconv.Atoi(texts[0])
	if err != nil {
		countOfRescales = 1
	}

	srcImageBuffers, err := utils.GetImages(obj.Message, 1)
	if err != nil {
		// Todo : ошибка получения изображения
		handler.Sender.Send(obj.Message.PeerID, err.Error())
		log.Println("[ERROR]: ", err)
		return
	}
	// image.Decode(...) съедает байты с буффера
	srcImageBytes := srcImageBuffers[0].Bytes()
	srcImage, _, err := image.Decode(srcImageBuffers[0])
	if err != nil {
		// Todo : ошибка unknown
		//handler.Sender.Send(obj.Message.PeerID, err.Error())
		log.Println("[ERROR]: ", err)
		return
	}

	mw := imagick.NewMagickWand()
	if err = mw.ReadImageBlob(srcImageBytes); err != nil {
		// Todo : ошибка unknown
		//handler.Sender.Send(obj.Message.PeerID, err.Error())
		log.Println("[ERROR]: ", err)
		return
	}

	for i := 1; i-1 < countOfRescales; i++ {
		if err = mw.LiquidRescaleImage(
			uint(float64(srcImage.Bounds().Size().X)/(1.2*float64(i))),
			uint(float64(srcImage.Bounds().Size().Y)/(1.2*float64(i))),
			0, 0,
		); err != nil {
			// Todo : ошибка unknown
			//handler.Sender.Send(obj.Message.PeerID, err.Error())
			log.Println("[ERROR]: ", err)
			return
		}
	}
	if err = mw.ResizeImage(
		uint(srcImage.Bounds().Size().X), uint(srcImage.Bounds().Size().X),
		imagick.FILTER_LANCZOS2,
	); err != nil {
		// Todo : ошибка unknown
		//handler.Sender.Send(obj.Message.PeerID, err.Error())
		log.Println("[ERROR]: ", err)
		return
	}

	outImageReader := utils.BytesToReader(mw.GetImageBlob())
	if err = handler.Sender.SendWithImage(obj.Message.PeerID, "Держи", outImageReader); err != nil {
		// Todo : ошибка отправки сообщения с изображением
		handler.Sender.Send(obj.Message.PeerID, err.Error())
		log.Println("[ERROR]: ", err)
		return
	}
	return
}
