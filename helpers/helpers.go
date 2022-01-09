package helpers

import (
	TYPES "gorth/types"
	"strings"
)

func FindInString(str, sub string, start int, end int) int {
	if start > len(str) || !strings.Contains(str[start:end], sub) {
		return -1
	}
	return strings.Index(str[start:end], sub) + start
}

func Enumerate(a []interface{}) []TYPES.Enumerator {
	enum := make([]TYPES.Enumerator, 0)
	for index, content := range a {
		enum = append(enum, TYPES.Enumerator{
			Index: uint(index + 1),
			Slice: content,
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

func EnumerateLine(line string, enumeration chan<- TYPES.Enumerator) {
	line = strings.TrimRight(line, "\n")
	col := findCol(line, 0, func(s string) bool {
		return s != " "
	})
	for col < len(line) {
		colEnd := findCol(line, col, func(s string) bool {
			return s == " "
		})
		enumeration <- TYPES.Enumerator{
			Index: uint(col),
			Slice: line[col:colEnd],
		}
		col = findCol(line, colEnd, func(s string) bool {
			return s != " "
		})
	}
	close(enumeration)
}
