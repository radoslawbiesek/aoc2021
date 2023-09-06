package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetLines(path, separator string) []string {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	content := string(data)
	lines := strings.Split(content, separator)

	return lines
}

func ParseInt(str string) int {
	parsed, err := strconv.Atoi(strings.TrimSpace(str))

	if err != nil {
		panic(fmt.Sprintf("Could not parse %s", str))
	}

	return parsed
}

func CharAt(str string, index int) string {
	return fmt.Sprintf("%c", str[index])
}

func IsNonEmptyStr(str string) bool {
	return strings.TrimSpace(str) != ""
}
