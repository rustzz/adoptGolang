package helpers

func IsDem(command string) bool {
	return command == "dem"
}

func IsTBD(command string) bool {
	return command == "tbd"
}

func IsСumCas(command string) bool {
	return command == "tbd" || command == "кас"
}
