package message

func IsDem(command string) bool {
	return command == "dem" || command == "дем" || command == "д" || command == "d"
}

func IsTBD(command string) bool {
	return command == "tbd" || command == "t"
}

func IsLiquidRescale(command string) bool {
	return command == "cum" || command == "кас" || command == "к"  || command == "c"
}
