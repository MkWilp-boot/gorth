package helpers

import (
	"gorth/types"
	"strings"
)

func FindInString(str, sub string, start int, end int) int {
	if start > len(str) || !strings.Contains(str[start:end], sub) {
		return -1
	}
	return strings.Index(str[start:end], sub) + start
}

func Enumerate(sVec2d []types.Vec2DString) []types.Vec2DString {
	enum := make([]types.Vec2DString, 0)

	for index, content := range sVec2d {
		enum = append(enum, types.Vec2DString{
			Index:   index + 1,
			Content: content.Content,
		})
	}
	return enum
}

func findCol(line string, start int, predicate func(string) bool) int {
	for start < len(line) && !predicate(string(line[start])) {
		start++
	}
	return start
}

func EnumerateLine(line string, enumeration chan<- types.StringEnum) {
	// checking for comments
	if strings.Contains(line, "//") {
		commentIndex := strings.Index(line, "//")
		line = line[:commentIndex]
	}

	line = strings.TrimSpace(line)

	col := findCol(line, 0, func(s string) bool {
		return s != " "
	})
	for col < len(line) {
		colEnd := findCol(line, col, func(s string) bool {
			return s == " "
		})
		enumeration <- types.StringEnum{
			Index: col,
			Slice: line[col:colEnd],
		}
		col = findCol(line, colEnd, func(s string) bool {
			return s != " "
		})
	}
	close(enumeration)
}
