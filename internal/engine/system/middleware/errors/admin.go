package errors

func (err *NotAdminError) Error() string {
	return "Вы не администратор."
}
func (err *BotNotAdminError) Error() string {
	return "Установите боту статус администратора."
}
