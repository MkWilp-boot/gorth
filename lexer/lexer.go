package lexer

import (
	HELP "gorth/helpers"
	TYPES "gorth/types"
	"io/ioutil"
	"strings"
)

func LexFile(filePath string) []TYPES.Enumerator {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	lines := make(TYPES.Enumerable, 0)
	for lineNumber, line := range strings.Split(string(bytes), "\n") {
		if len(line) != 0 {
			enumeration := make(chan TYPES.Enumerator)
			go HELP.EnumerateLine(line, enumeration)
			for enumeratedLine := range enumeration {
				lines = append(lines, TYPES.Enumerator{
					Index: uint(lineNumber + 1),
					Slice: enumeratedLine,
				})
			}
		}
	}
	return HELP.Enumerate(lines)
}
