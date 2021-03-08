package checker

import (
	"adoptGolang/internal/handler/adopt"
	"regexp"
	"strings"
)

func IsBotNamePrefix(message string) (matched bool) {
	_tmp := strings.ToLower(strings.Split(message, " ")[0])
	matched, _ = regexp.Match(adopt.Prefix, []byte(_tmp))
	return
}
