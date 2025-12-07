package util

import (
	"bufio"
	"bytes"
	"embed"
	"strings"
)

func LoadFile(filename string, fileSystem embed.FS) ([]string, error) {
	bytesLines, err := fileSystem.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	stringLines := strings.TrimSpace(string(bytesLines))

	return strings.Split(stringLines, "\n"), nil
}

func LoadCommaFile(filename string, fileSystem embed.FS) ([]string, error) {
	bytesLines, err := fileSystem.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(bytesLines)), ","), nil
}

func LoadSplitFile(filename string, fileSystem embed.FS) ([]string, []string, error) {
	data, err := fileSystem.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}

	var ranges []string
	var numbers []string

	scanner := bufio.NewScanner(bytes.NewReader(data))
	section := 0 // 0 = ranges, 1 = numbers

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			section = 1
			continue
		}

		if section == 0 {
			ranges = append(ranges, line)
		} else {
			numbers = append(numbers, line)
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, nil, err
	}

	return ranges, numbers, nil
}
