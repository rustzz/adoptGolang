package errors

func (err *ImageNotFound) Error() string {
	return "Изображение(-я) не найдено(-ы)"
}
