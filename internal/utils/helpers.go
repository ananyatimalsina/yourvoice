package utils

import (
	"strings"
)

func Idfy(str string) string {
	return strings.ToLower(strings.ReplaceAll(str, " ", "-"))
}
