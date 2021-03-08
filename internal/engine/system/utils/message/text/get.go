package text

import (
	"adoptGolang/internal/engine/system/checker"
	"fmt"
	"regexp"
	"strings"
)

const (
	mentionPattern = `^\[\w+\d+\|.*\w+\].*$`
)

/*
GetCommandArguments : Отдает аргументы команды без имени команды и упоминания
*/
func GetCommandArguments(message string, reqCountOfLines int) (out []string) {
	messageLines := strings.Split(message, "\n")

	var fromIndex int /*
		от какого индекса получать текста команды
		ниже: если имеется префикс к боту в виде [club123456|name],
		то отдавать от первого индекса, где находится имя команды
	*/
	matched, _ := regexp.Match(
		mentionPattern, []byte(strings.ToLower(strings.Split(message, " ")[0])))
	if matched || checker.IsBotNamePrefix(message) {
		fromIndex = 1
	}

	// получить текст с первой строки от имени команды
	var _tmp string
	for index, word := range strings.Split(messageLines[0], " ") {
		if index > fromIndex {
			_tmp += fmt.Sprintf("%s ", word)
		}
	}

	// добавить строки
	if len(_tmp) > 0 {
		out = append(out, _tmp[0:len(_tmp)-1])
	}
	for _, elem := range messageLines[1:] {
		out = append(out, elem)
	}
	// не достающие строки заполнить пустотой воизбежание нехватки индексов
	if len(out) != reqCountOfLines {
		for i := len(out); i < reqCountOfLines; i++ {
			out = append(out, "")
		}
	}
	return
}

/*
GetCommand : Отдает имя команды в зависимости как вызывался бот
			[club123456|name], -  с запятой для моб. клиента
								  без запятой для веб клиента
*/
func GetCommand(message string) string {
	matched, _ := regexp.Match(
		mentionPattern, []byte(strings.ToLower(strings.Split(message, " ")[0])))
	if matched || checker.IsBotNamePrefix(message) {
		if len(strings.Split(message, " ")) < 2 { return "" }
		return strings.ToLower(strings.Split(message, " ")[1])
	}
	return strings.ToLower(strings.Split(message, " ")[0])
}
