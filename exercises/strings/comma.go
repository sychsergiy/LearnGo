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

func splitOnChunksFromStart(s string, chunkSize int) []string {
	// for IntegerPart
	var chunks []string
	sLen := len(s)
	rest := sLen % chunkSize
	if rest > 0 {
		chunks = append(chunks, s[:rest])
	}
	chunks = append(chunks, splitOnChunks(s[rest:], chunkSize)...)
	return chunks
}

func splitOnChunksFromEnd(s string, chunkSize int) []string {
	// for Float part
	sLen := len(s)
	rest := sLen % chunkSize
	chunks := splitOnChunks(s[:sLen-rest], chunkSize)
	if rest > 0 {
		chunks = append(chunks, s[sLen-rest:])
	}
	return chunks
}

func splitOnChunks(s string, chunkSize int) []string {
	var chunks []string

	sLen := len(s)
	if sLen%chunkSize != 0 {
		panic("Wrong input argument")
	}
	for i := 0; i < len(s); i += chunkSize {
		chunks = append(chunks, s[i:i+chunkSize])
	}
	return chunks
}

func insertCommasToIntegerPart(s string, chunkSize int) string {
	chunks := splitOnChunksFromStart(s, chunkSize)
	return strings.Join(chunks, ",")
}

func insertCommasToFloatPart(s string, chunkSize int) string {
	chunks := splitOnChunksFromEnd(s, chunkSize)
	return strings.Join(chunks, ",")
}

func insertCommas(s string, chunkSize int) string {
	chunks := strings.Split(s, ".")
	if len(chunks) == 1 {
		intPart := chunks[0]
		return insertCommasToIntegerPart(intPart, chunkSize)
	}
	if len(chunks) == 2 {
		intPart, floatPart := chunks[0], chunks[1]
		intPartWithCommas := insertCommasToIntegerPart(intPart, chunkSize)
		floatPartWithCommas := insertCommasToFloatPart(floatPart, chunkSize)
		return intPartWithCommas + "." + floatPartWithCommas

	} else {
		panic("Wrong input")
	}
}
