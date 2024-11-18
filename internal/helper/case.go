package helper

import (
	"github.com/Pashgunt/converter/internal/enum"
	"regexp"
	"strings"
)

func CamelToSnake(input string) string {
	return strings.ToLower(regexp.MustCompile(enum.CamelCase).ReplaceAllString(input, enum.SnackCase))
}
