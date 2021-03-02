package text

func IsDem(command string) bool {
	return command == "dem" || command == "дем"
}

func IsTBD(command string) bool {
	return command == "tbd" || command == "t"
}

func IsLiquidRescale(command string) bool {
	return command == "cum" || command == "кас"
}
