package common

import (
	"os"
	"strings"
)

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFile(path string) []string {
	content, err := os.ReadFile(path)
	PanicIfErr(err)
	return strings.Split(strings.TrimSpace(string(content)), "\n")
}
