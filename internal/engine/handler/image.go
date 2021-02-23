package handler

import (
	"adoptGolang/internal/engine"
	"adoptGolang/internal/engine/errors"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/rustzz/blocks"
	"github.com/rustzz/demotivator"
	"image"
	"log"
)

// HandleDem : Демотиватор
func (handler *Handler) HandleDem(obj *events.MessageNewObject) {
	dem := demotivator.New()
	srcImageReaders, err := engine.GetImages(obj.Message, 1)
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
	texts := engine.GetTexts(obj.Message.Text, 2)
	outImageReader, err := dem.Make(&srcImage, texts, "")
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

// HandleTBD : ...
func (handler *Handler) HandleTBD(obj *events.MessageNewObject) {
	tbd := blocks.New()
	srcImageReaders, err := engine.GetImages(obj.Message, 2)
	if err != nil {
		handler.Sender.Send(obj.Message.PeerID, err.Error())
		log.Println("[ERROR]: ", err)
		return
	}
	images, _ := func () (out []*image.Image, err error) {
		for _, srcImageReader := range srcImageReaders {
			srcImage, _, err := image.Decode(srcImageReader)
			if err != nil {
				handler.Sender.Send(obj.Message.PeerID, err.Error())
				log.Println("[ERROR]: ", err)
				return out, err
			}
			out = append(out, &srcImage)
		}
		return
	}()
	texts := engine.GetTexts(obj.Message.Text, 3)
	outImageReader, err := tbd.Make(&images, texts, "")
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

// HandleLiquidRescale : функция кас
func (handler *Handler) HandleLiquidRescale(obj *events.MessageNewObject) {
	handler.Sender.Send(obj.Message.PeerID, new(errors.ModuleNotImplemented).Error())
	return
}
