package handler

import (
	"adoptGolang/internal/engine/adopt/errors"
	"adoptGolang/internal/engine/system/utils"
	"github.com/SevereCloud/vksdk/v2/events"
	mBlocks "github.com/rustzz/blocks"
	mDemotivator "github.com/rustzz/demotivator"
	"image"
	"log"
)

// HandleDem : Демотиватор
func (handler *AdoptHandler) HandleDem(obj events.MessageNewObject) {
	texts := utils.GetTexts(obj.Message.Text, 2)

	srcImageBuffers, err := utils.GetImages(obj.Message, 1)
	if err != nil {
		handler.Handler.Sender.Send(obj.Message.PeerID, new(errors.GetImagesError).Error())
		log.Println("[ERROR]: ", err)
		return
	}
	srcImage, _, err := image.Decode(srcImageBuffers[0])
	if err != nil {
		handler.Handler.Sender.Send(obj.Message.PeerID, new(errors.UnknownError).Error())
		log.Println("[ERROR]: ", err)
		return
	}

	demotivator := mDemotivator.New(srcImage, [2]string{texts[0], texts[1]})
	outImageReader, err := demotivator.Make(nil, [2]string{})
	if err != nil {
		handler.Handler.Sender.Send(obj.Message.PeerID, new(errors.UnknownError).Error())
		log.Println("[ERROR]: ", err)
		return
	}
	handler.Handler.Sender.SendWithImage(
		obj.Message.PeerID, "Держи", utils.BufferToReader(outImageReader))
	return
}

// HandleTBD : ...
func (handler *AdoptHandler) HandleTBD(obj events.MessageNewObject) {
	texts := utils.GetTexts(obj.Message.Text, 3)

	srcImageBuffers, err := utils.GetImages(obj.Message, 2)
	if err != nil {
		handler.Handler.Sender.Send(obj.Message.PeerID, new(errors.GetImagesError).Error())
		log.Println("[ERROR]: ", err)
		return
	}
	images, _ := func () (out []image.Image, err error) {
		for _, srcImageBuffer := range srcImageBuffers {
			srcImage, _, err := image.Decode(srcImageBuffer)
			if err != nil {
				handler.Handler.Sender.Send(obj.Message.PeerID, new(errors.UnknownError).Error())
				log.Println("[ERROR]: ", err)
				return nil, err
			}
			out = append(out, srcImage)
		}
		return
	}()

	blocks := mBlocks.New(
		[2]image.Image{images[0], images[1]},
		[3]string{texts[0], texts[1], texts[2]},
	)
	outImageReader, err := blocks.Make()
	if err != nil {
		handler.Handler.Sender.Send(obj.Message.PeerID, new(errors.UnknownError).Error())
		log.Println("[ERROR]: ", err)
		return
	}
	handler.Handler.Sender.SendWithImage(
		obj.Message.PeerID, "Держи", utils.BufferToReader(outImageReader))
	return
}

// HandleLiquidRescale : функция кас
func (handler *AdoptHandler) HandleLiquidRescale(obj events.MessageNewObject) {
	// Todo : Смотреть внутрь функции
	DoRescale(handler, obj)
	return
}
