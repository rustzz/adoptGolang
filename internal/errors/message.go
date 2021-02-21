package errors

func (err *FwdReplyMessageNotFound) Error() string {
	return "Нет изображения, ответного сообщения, пересланного сообщения"
}
