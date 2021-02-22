package handler

import (
	"github.com/SevereCloud/vksdk/v2/events"
	"adoptGolang/internal/helpers"
	"github.com/rustzz/blocks"
	"github.com/rustzz/demotivator"
	"image"
	"log"
)

func (handler *Handler) HandleDem(obj *events.MessageNewObject) {
	dem := &demotivator.Demotivator{}
	srcImageReaders, err := helpers.GetImages(obj.Message, 1)
	if err != nil {
		handler.Sender.Send(obj.Message.PeerID, err.Error())
		log.Println("[ERROR]: ", err)
		return
	}
	srcImage, _, err := image.Decode(srcImageReaders[0])
	if err != nil {
		handler.Sender.Send(obj.Message.PeerID, err.Error())
		log.Println("[ERROR]: ", err)
		return
	}
	texts := helpers.GetTexts(obj.Message.Text, 2)
	outImageReader, err := dem.Make(
		srcImage, texts, "",
	)
	if err != nil {
		handler.Sender.Send(obj.Message.PeerID, err.Error())
		log.Println("[ERROR]: ", err)
		return
	}
	if err = handler.Sender.SendWithImage(obj.Message.PeerID, "Держи", outImageReader); err != nil {
		handler.Sender.Send(obj.Message.PeerID, err.Error())
		log.Println("[ERROR]: ", err)
		return
	}
	return
}

func (handler *Handler) HandleTBD(obj *events.MessageNewObject) {
	tbd := blocks.New()
	srcImageReaders, err := helpers.GetImages(obj.Message, 2)
	if err != nil {
		handler.Sender.Send(obj.Message.PeerID, err.Error())
		log.Println("[ERROR]: ", err)
		return
	}
	images := func () (out []image.Image) {
		for _, srcImageReader := range srcImageReaders {
			srcImage, _, err := image.Decode(srcImageReader)
			if err != nil {
				handler.Sender.Send(obj.Message.PeerID, err.Error())
				log.Println("[ERROR]: ", err)
				return
			}
			out = append(out, srcImage)
		}
		return
	}()
	texts := helpers.GetTexts(obj.Message.Text, 3)
	outImageReader, err := tbd.Make(
		images, texts, "",
	)
	if err != nil {
		handler.Sender.Send(obj.Message.PeerID, err.Error())
		log.Println("[ERROR]: ", err)
		return
	}
	if err = handler.Sender.SendWithImage(obj.Message.PeerID, "Держи", outImageReader); err != nil {
		handler.Sender.Send(obj.Message.PeerID, err.Error())
		log.Println("[ERROR]: ", err)
		return
	}
	return
}
