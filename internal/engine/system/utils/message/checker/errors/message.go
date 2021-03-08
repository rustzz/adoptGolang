package errors

func (err *FwdMessageNotFound) Error() string {
	return "Нет пересланного сообщения."
}

func (err *RepliedMessageNotFound) Error() string {
	return "Нет отвеченного сообщения."
}
