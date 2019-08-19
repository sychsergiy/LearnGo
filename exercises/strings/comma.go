package strings

import (
	"strings"
)

func commaRecursive(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaRecursive(s[:n-3]) + "," + s[n-3:]
}

func commaLoop(s string) string {
	var chunks []string

	sLen := len(s)
	offset := sLen % 3

	if offset > 0 {
		chunks = append(chunks, s[:offset])
	}
	for i := offset; i < sLen; i += 3 {
		chunks = append(chunks, s[i:i+3])
	}

	return strings.Join(chunks, ",")
}

func commaFloatNumber(s string) string {
	chunks := strings.Split(s, ",")
	if len(chunks) > 2 {
		panic("Wrong input")
	}
	intPart, floatPart := chunks[0], chunks[1]

	intPartWithCommas := commaLoop(intPart)
	floatPartWithCommas := commaLoop(floatPart)
	return intPartWithCommas + "." + floatPartWithCommas
}
