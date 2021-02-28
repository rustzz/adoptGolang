package handler

import (
	"adoptGolang/internal/engine/utils"
	"adoptGolang/pkg/rescale"
	"github.com/SevereCloud/vksdk/v2/events"
	mBlocks "github.com/rustzz/blocks"
	mDemotivator "github.com/rustzz/demotivator"
	"image"
	"log"
)

// HandleDem : Демотиватор
func (handler *Handler) HandleDem(obj events.MessageNewObject) {
	texts := utils.GetTexts(obj.Message.Text, 2)

	srcImageBuffers, err := utils.GetImages(obj.Message, 1)
	if err != nil {
		// Todo : Создать кастомную ошибку получения изображений
		handler.Sender.Send(obj.Message.PeerID, err.Error())
		log.Println("[ERROR]: ", err)
		return
	}
	srcImage, _, err := image.Decode(srcImageBuffers[0])
	if err != nil {
		// Todo : Создать кастомную unknown ошибку, т.к. это нельзя показывать пользователю
		//handler.Sender.Send(obj.Message.PeerID, err.Error())
		log.Println("[ERROR]: ", err)
		return
	}

	demotivator := mDemotivator.New(srcImage, [2]string{texts[0], texts[1]})
	outImageReader, err := demotivator.Make(nil, [2]string{})
	if err != nil {
		// Todo : ошибка unknown
		//handler.Sender.Send(obj.Message.PeerID, err.Error())
		log.Println("[ERROR]: ", err)
		return
	}
	if err = handler.Sender.SendWithImage(
		obj.Message.PeerID, "Держи", utils.BufferToReader(outImageReader),
	); err != nil {
		// Todo : Создать кастомную ошибку отправки сообщения с изображением
		handler.Sender.Send(obj.Message.PeerID, err.Error())
		log.Println("[ERROR]: ", err)
		return
	}
	return
}

// HandleTBD : ...
func (handler *Handler) HandleTBD(obj events.MessageNewObject) {
	texts := utils.GetTexts(obj.Message.Text, 3)

	srcImageBuffers, err := utils.GetImages(obj.Message, 2)
	if err != nil {
		// Todo : ошибка получения изображения
		handler.Sender.Send(obj.Message.PeerID, err.Error())
		log.Println("[ERROR]: ", err)
		return
	}
	images, _ := func () (out []image.Image, err error) {
		for _, srcImageBuffer := range srcImageBuffers {
			srcImage, _, err := image.Decode(srcImageBuffer)
			if err != nil {
				// Todo : ошибка unknown
				//handler.Sender.Send(obj.Message.PeerID, err.Error())
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
		// Todo : ошибка unknown
		//handler.Sender.Send(obj.Message.PeerID, err.Error())
		log.Println("[ERROR]: ", err)
		return
	}
	if err = handler.Sender.SendWithImage(
		obj.Message.PeerID, "Держи", utils.BufferToReader(outImageReader),
	); err != nil {
		// Todo : ошибка отправки сообщения с изображением
		handler.Sender.Send(obj.Message.PeerID, err.Error())
		log.Println("[ERROR]: ", err)
		return
	}
	return
}

// HandleLiquidRescale : функция кас
func (handler *Handler) HandleLiquidRescale(obj events.MessageNewObject) {
	// Todo : Смотреть внутрь функции
	rescale.DoRescale(handler, obj)
	return
}
