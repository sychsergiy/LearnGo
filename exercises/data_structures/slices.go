package data_structures

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseArray(sPtr *[]int) {
	s := *sPtr
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, n int) {
	// Rotate s left by two positions.
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

func rotateByCopy(s []int, n int) {
	sCopy := make([]int, n)
	copy(sCopy, s[:n])          // copy before n to another slice
	copy(s, s[n:])              // copy after n to the begin of slice
	copy(s[len(s)-n:], sCopy, ) // copy before n from another slice to the end of current
}

func rotate3(s []int, n int) {
	// use pop and push
}
