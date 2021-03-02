package errors

func (err *ImageNotFound) Error() string {
	return "Изображение(-я) не найдено(-ы)"
}

func (err *GetImagesError) Error() string {
	return "Не удалось получить изображения из сообщения"
}
