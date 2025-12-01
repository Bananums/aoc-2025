package util

import (
	"bufio"
	"embed"
	"os"
	"strings"
)

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func LoadFile(filename string, fileSystem embed.FS) ([]string, error) {
	bytes, err := fileSystem.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return strings.Split(strings.TrimSpace(string(bytes)), "\n"), nil
}
