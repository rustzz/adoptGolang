package handler

import (
	"adoptGolang/internal/engine/adopt/errors"
	"adoptGolang/internal/engine/system/utils"
	"github.com/SevereCloud/vksdk/v2/events"
	"gopkg.in/gographics/imagick.v3/imagick"
	"image"
	"log"
	"strconv"
)

// Todo : Перенести в внешний модуль (репозиторий)
func DoRescale(handler *AdoptHandler, obj events.MessageNewObject) {
	imagick.Initialize()
	defer imagick.Terminate()

	texts := utils.GetTexts(obj.Message.Text, 1)
	log.Println(texts)
	countOfRescales, err := strconv.Atoi(texts[0])
	if err != nil {
		countOfRescales = 1
	}

	srcImageBuffers, err := utils.GetImages(obj.Message, 1)
	if err != nil {
		handler.Handler.Sender.Send(obj.Message.PeerID, new(errors.GetImagesError).Error())
		log.Println("[ERROR]: ", err)
		return
	}
	// image.Decode(...) съедает байты с буффера
	srcImageBytes := srcImageBuffers[0].Bytes()
	srcImage, _, err := image.Decode(srcImageBuffers[0])
	if err != nil {
		handler.Handler.Sender.Send(obj.Message.PeerID, new(errors.UnknownError).Error())
		log.Println("[ERROR]: ", err)
		return
	}

	mw := imagick.NewMagickWand()
	if err = mw.ReadImageBlob(srcImageBytes); err != nil {
		handler.Handler.Sender.Send(obj.Message.PeerID, new(errors.UnknownError).Error())
		log.Println("[ERROR]: ", err)
		return
	}

	for i := 1; i-1 < countOfRescales; i++ {
		if err = mw.LiquidRescaleImage(
			uint(float64(srcImage.Bounds().Size().X)/(1.2*float64(i))),
			uint(float64(srcImage.Bounds().Size().Y)/(1.2*float64(i))),
			0, 0,
		); err != nil {
			handler.Handler.Sender.Send(obj.Message.PeerID, new(errors.UnknownError).Error())
			log.Println("[ERROR]: ", err)
			return
		}
	}
	if err = mw.ResizeImage(
		uint(srcImage.Bounds().Size().X), uint(srcImage.Bounds().Size().X),
		imagick.FILTER_LANCZOS2,
	); err != nil {
		handler.Handler.Sender.Send(obj.Message.PeerID, new(errors.UnknownError).Error())
		log.Println("[ERROR]: ", err)
		return
	}

	outImageReader := utils.BytesToReader(mw.GetImageBlob())
	handler.Handler.Sender.SendWithImage(obj.Message.PeerID, "Держи", outImageReader)
	return
}
