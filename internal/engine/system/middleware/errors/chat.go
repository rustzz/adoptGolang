package errors

func (err *IsNotChatError) Error() string {
	return "Используйте в чате."
}

func (err *IsNotPrivateChatError) Error() string {
	return "Используйте в ЛС."
}
