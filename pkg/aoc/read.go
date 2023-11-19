package aoc

import (
	"bufio"
	"io"
)

func ParseLines[T any](reader io.Reader, fn func(string) (T, error)) ([]T, error) {
	scanner := bufio.NewScanner(reader)
	var result []T
	for scanner.Scan() {
		line := scanner.Text()
		value, err := fn(line)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, scanner.Err()
}

func ReadLines(reader io.Reader) ([]string, error) {
	return ParseLines(reader, func(line string) (string, error) {
		return line, nil
	})
}
