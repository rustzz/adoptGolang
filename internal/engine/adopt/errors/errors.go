package errors

type ImageNotFound struct {}
type GetImagesError struct {}
type FwdReplyMessageNotFound struct {}
type ModuleNotImplemented struct {}
type UnknownError struct {}

func (err *UnknownError) Error() string {
	return "Неизвестная ошибка"
}
