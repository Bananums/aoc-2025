package util

import (
	"embed"
	"strings"
)

func LoadFile(filename string, fileSystem embed.FS) ([]string, error) {
	bytes, err := fileSystem.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return strings.Split(strings.TrimSpace(string(bytes)), "\n"), nil
}

func LoadCommaFile(filename string, fileSystem embed.FS) ([]string, error) {
	bytes, err := fileSystem.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(bytes)), ","), nil
}
